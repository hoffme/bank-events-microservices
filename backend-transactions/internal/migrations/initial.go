package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func run_initial(ctx context.Context, client *mongo.Client) error {
	var err error

	db := client.Database("bank_transactions")

	// views

	pipeline := mongo.Pipeline{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "accounts"},
					{"localField", "from_account_id"},
					{"foreignField", "_id"},
					{"as", "from_account"},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "accounts"},
					{"localField", "to_account_id"},
					{"foreignField", "_id"},
					{"as", "to_account"},
				},
			},
		},
		bson.D{{"$unwind", "$from_account"}},
		bson.D{{"$unwind", "$to_account"}},
		bson.D{
			{"$project",
				bson.D{
					{"_id", "$_id"},
					{"from_account_id", "$from_account._id"},
					{"from_account_name", "$from_account.name"},
					{"to_account_id", "$to_account._id"},
					{"to_account_name", "$to_account.name"},
					{"state", "$state"},
					{"amount", "$amount"},
					{"currency", "$currency"},
					{"created_at", "$created_at"},
					{"finished_at", "$finished_at"},
				},
			},
		},
	}

	err = db.Collection("view_transactions").Drop(ctx)
	if err != nil {
		return err
	}

	err = db.CreateView(ctx, "view_transactions", "transactions", pipeline)
	if err != nil {
		return err
	}

	return nil
}
