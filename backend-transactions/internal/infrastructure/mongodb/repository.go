package mongodb

import (
	"context"
	"errors"

	"github.com/hoffme/backend-transactions/internal/shared/null"
	"github.com/hoffme/backend-transactions/internal/shared/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo[R any] struct {
	connection       *Connection
	databaseName     string
	collectionName   string
	notFoundError    error
	fieldsProperties map[string]string
}

func (u repo[R]) get(ctx context.Context, id string) (R, error) {
	doc := new(R)

	mongoFilter := bson.M{"_id": id}

	err := u.collection().FindOne(ctx, mongoFilter).Decode(doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return *doc, u.notFoundError
	}
	if err != nil {
		return *doc, err
	}

	return *doc, nil
}

func (u repo[R]) find(ctx context.Context, filter repository.FindFilter) (R, error) {
	doc := new(R)

	mongoFilter, mongoOptions := u.encodeFindFilter(filter)

	err := u.collection().FindOne(ctx, mongoFilter, mongoOptions).Decode(doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return *doc, u.notFoundError
	}
	if err != nil {
		return *doc, err
	}

	return *doc, nil
}

func (u repo[R]) search(ctx context.Context, filter repository.SearchFilter) (repository.SearchResult[R], error) {
	filterMongo, optionsMongo := u.encodeSearchFilter(filter)

	cursor, err := u.collection().Find(ctx, filterMongo, optionsMongo)
	if err != nil {
		return repository.SearchResult[R]{}, err
	}

	count, err := u.collection().CountDocuments(ctx, filterMongo)
	if err != nil {
		return repository.SearchResult[R]{}, err
	}

	result := repository.SearchResult[R]{
		Data:  []R{},
		Count: count,
	}

	if optionsMongo.Skip != nil {
		result.Skip = *optionsMongo.Skip
	} else {
		result.Skip = 0
	}

	if optionsMongo.Limit != nil {
		result.Limit = *optionsMongo.Limit
	} else {
		result.Limit = 0
	}

	err = cursor.All(ctx, &result.Data)
	if err != nil {
		return repository.SearchResult[R]{}, err
	}

	return result, nil
}

func (u repo[R]) save(ctx context.Context, id string, record R) (bool, error) {
	mongoFilter := bson.M{"_id": id}
	mongoUpdate := bson.M{"$set": record}
	mongoOptions := options.Update().SetUpsert(true)

	result, err := u.collection().UpdateOne(ctx, mongoFilter, mongoUpdate, mongoOptions)
	if err != nil {
		return false, err
	}

	return result.ModifiedCount > 0, nil
}

func (u repo[R]) delete(ctx context.Context, id string) (bool, error) {
	mongoFilter := bson.M{"_id": id}

	result, err := u.collection().DeleteOne(ctx, mongoFilter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}

// utils

func (u repo[R]) collection() *mongo.Collection {
	return u.connection.client.Database(u.databaseName).Collection(u.collectionName)
}

func (u repo[R]) encodeFindFilter(filter repository.FindFilter) (bson.M, *options.FindOneOptions) {
	mongoOptions := options.FindOne()

	if filter.Order.IsNotNull() {
		for _, order := range filter.Order.Value() {
			mongoKey, ok := u.fieldsProperties[order.By]
			if !ok {
				continue
			}

			dir := 1
			if order.Dir == repository.OrderDirDesc {
				dir = -1
			}

			mongoOptions = mongoOptions.SetSort(bson.M{mongoKey: dir})
		}
	}

	return u.encodeFilter(filter.Query, filter.Where), mongoOptions
}

func (u repo[R]) encodeSearchFilter(filter repository.SearchFilter) (bson.M, *options.FindOptions) {
	mongoOptions := options.Find()

	if filter.Order.IsNotNull() {
		for _, order := range filter.Order.Value() {
			mongoKey, ok := u.fieldsProperties[order.By]
			if !ok {
				continue
			}

			dir := 1
			if order.Dir == repository.OrderDirDesc {
				dir = -1
			}

			mongoOptions = mongoOptions.SetSort(bson.M{mongoKey: dir})
		}
	}

	if filter.Skip.IsNotNull() {
		mongoOptions = mongoOptions.SetSkip(filter.Skip.Value())
	}

	if filter.Limit.IsNotNull() {
		mongoOptions = mongoOptions.SetLimit(filter.Limit.Value())
	}

	return u.encodeFilter(filter.Query, filter.Where), mongoOptions
}

func (u repo[R]) encodeFilter(query null.Null[string], where null.Null[repository.Where]) bson.M {
	mongoFilter := bson.A{}

	if query.IsNotNull() {
		mongoFilter = append(mongoFilter, bson.M{"$text": bson.M{"$search": query.Value()}})
	}

	if where.IsNotNull() {
		subQuery, ok := u.decodeWhere(where.Value())
		if ok {
			mongoFilter = append(mongoFilter, subQuery)
		}
	}

	if len(mongoFilter) > 0 {
		return bson.M{"$and": mongoFilter}
	}

	return bson.M{}
}

func (u repo[R]) decodeWhere(where repository.Where) (bson.M, bool) {
	if len(where.Field) > 0 {
		location, ok := u.fieldsProperties[where.Field]
		if !ok {
			return bson.M{}, false
		}

		switch where.Operator {
		case repository.WhereOpEq:
			return bson.M{location: bson.M{"$eq": where.Value}}, true
		case repository.WhereOpLte:
			return bson.M{location: bson.M{"$lte": where.Value}}, true
		case repository.WhereOpLt:
			return bson.M{location: bson.M{"$lt": where.Value}}, true
		case repository.WhereOpGte:
			return bson.M{location: bson.M{"$gte": where.Value}}, true
		case repository.WhereOpGt:
			return bson.M{location: bson.M{"$gte": where.Value}}, true
		case repository.WhereOpIn:
			return bson.M{location: bson.M{"$in": where.Value}}, true
		case repository.WhereOpRgx:
			return bson.M{location: where.Value}, true
		case repository.WhereOpBtw:
			values, ok := where.Value.([]interface{})
			if !ok || len(values) != 2 {
				return bson.M{}, false
			}
			return bson.M{location: bson.M{"$gte": values[0], "$lt": values[1]}}, true
		case repository.WhereOpNot:
			subWhere, ok := where.Value.(repository.Where)
			if !ok {
				return bson.M{}, false
			}

			subQuery, ok := u.decodeWhere(subWhere)
			if !ok {
				return bson.M{}, false
			}

			return bson.M{location: bson.M{"$not": subQuery}}, true
		case repository.WhereOpAnd:
			subWheres, ok := where.Value.([]repository.Where)
			if !ok {
				return bson.M{}, false
			}

			queryAnd := bson.A{}
			for _, subWhere := range subWheres {
				subQuery, ok := u.decodeWhere(subWhere)
				if ok {
					queryAnd = append(queryAnd, subQuery)
				}
			}
			if len(queryAnd) == 0 {
				return bson.M{}, false
			}

			return bson.M{location: bson.M{"$and": queryAnd}}, true
		case repository.WhereOpOr:
			subWheres, ok := where.Value.([]repository.Where)
			if !ok {
				return bson.M{}, false
			}

			queryOr := bson.A{}
			for _, subWhere := range subWheres {
				subQuery, ok := u.decodeWhere(subWhere)
				if ok {
					queryOr = append(queryOr, subQuery)
				}
			}
			if len(queryOr) == 0 {
				return bson.M{}, false
			}

			return bson.M{location: bson.M{"$or": queryOr}}, true
		}

		return bson.M{}, false
	}

	switch where.Operator {
	case repository.WhereOpNot:
		subWhere, ok := where.Value.(repository.Where)
		if !ok {
			return bson.M{}, false
		}

		subQuery, ok := u.decodeWhere(subWhere)
		if !ok {
			return bson.M{}, false
		}

		return bson.M{"$not": subQuery}, true
	case repository.WhereOpAnd:
		subWheres, ok := where.Value.([]repository.Where)
		if !ok {
			return bson.M{}, false
		}

		queryAnd := bson.A{}
		for _, subWhere := range subWheres {
			subQuery, ok := u.decodeWhere(subWhere)
			if ok {
				queryAnd = append(queryAnd, subQuery)
			}
		}
		if len(queryAnd) == 0 {
			return bson.M{}, false
		}

		return bson.M{"$and": queryAnd}, true
	case repository.WhereOpOr:
		subWheres, ok := where.Value.([]repository.Where)
		if !ok {
			return bson.M{}, false
		}

		queryOr := bson.A{}
		for _, subWhere := range subWheres {
			subQuery, ok := u.decodeWhere(subWhere)
			if ok {
				queryOr = append(queryOr, subQuery)
			}
		}
		if len(queryOr) == 0 {
			return bson.M{}, false
		}

		return bson.M{"$or": queryOr}, true
	}

	return bson.M{}, false
}
