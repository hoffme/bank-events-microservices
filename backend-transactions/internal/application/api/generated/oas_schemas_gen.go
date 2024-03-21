// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Currency to create Transaction.
// Ref: #/components/schemas/Currency
type Currency string

const (
	CurrencyARS Currency = "ARS"
	CurrencyUSD Currency = "USD"
)

// AllValues returns all Currency values.
func (Currency) AllValues() []Currency {
	return []Currency{
		CurrencyARS,
		CurrencyUSD,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Currency) MarshalText() ([]byte, error) {
	switch s {
	case CurrencyARS:
		return []byte(s), nil
	case CurrencyUSD:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Currency) UnmarshalText(data []byte) error {
	switch Currency(data) {
	case CurrencyARS:
		*s = CurrencyARS
		return nil
	case CurrencyUSD:
		*s = CurrencyUSD
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Status      int       `json:"status"`
	Code        string    `json:"code"`
	Description OptString `json:"description"`
}

// GetStatus returns the value of Status.
func (s *Error) GetStatus() int {
	return s.Status
}

// GetCode returns the value of Code.
func (s *Error) GetCode() string {
	return s.Code
}

// GetDescription returns the value of Description.
func (s *Error) GetDescription() OptString {
	return s.Description
}

// SetStatus sets the value of Status.
func (s *Error) SetStatus(val int) {
	s.Status = val
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val string) {
	s.Code = val
}

// SetDescription sets the value of Description.
func (s *Error) SetDescription(val OptString) {
	s.Description = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTransactionsGetOrderBy returns new OptTransactionsGetOrderBy with value set to v.
func NewOptTransactionsGetOrderBy(v TransactionsGetOrderBy) OptTransactionsGetOrderBy {
	return OptTransactionsGetOrderBy{
		Value: v,
		Set:   true,
	}
}

// OptTransactionsGetOrderBy is optional TransactionsGetOrderBy.
type OptTransactionsGetOrderBy struct {
	Value TransactionsGetOrderBy
	Set   bool
}

// IsSet returns true if OptTransactionsGetOrderBy was set.
func (o OptTransactionsGetOrderBy) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTransactionsGetOrderBy) Reset() {
	var v TransactionsGetOrderBy
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTransactionsGetOrderBy) SetTo(v TransactionsGetOrderBy) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTransactionsGetOrderBy) Get() (v TransactionsGetOrderBy, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTransactionsGetOrderBy) Or(d TransactionsGetOrderBy) TransactionsGetOrderBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTransactionsGetOrderDir returns new OptTransactionsGetOrderDir with value set to v.
func NewOptTransactionsGetOrderDir(v TransactionsGetOrderDir) OptTransactionsGetOrderDir {
	return OptTransactionsGetOrderDir{
		Value: v,
		Set:   true,
	}
}

// OptTransactionsGetOrderDir is optional TransactionsGetOrderDir.
type OptTransactionsGetOrderDir struct {
	Value TransactionsGetOrderDir
	Set   bool
}

// IsSet returns true if OptTransactionsGetOrderDir was set.
func (o OptTransactionsGetOrderDir) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTransactionsGetOrderDir) Reset() {
	var v TransactionsGetOrderDir
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTransactionsGetOrderDir) SetTo(v TransactionsGetOrderDir) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTransactionsGetOrderDir) Get() (v TransactionsGetOrderDir, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTransactionsGetOrderDir) Or(d TransactionsGetOrderDir) TransactionsGetOrderDir {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTransactionsGetType returns new OptTransactionsGetType with value set to v.
func NewOptTransactionsGetType(v TransactionsGetType) OptTransactionsGetType {
	return OptTransactionsGetType{
		Value: v,
		Set:   true,
	}
}

// OptTransactionsGetType is optional TransactionsGetType.
type OptTransactionsGetType struct {
	Value TransactionsGetType
	Set   bool
}

// IsSet returns true if OptTransactionsGetType was set.
func (o OptTransactionsGetType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTransactionsGetType) Reset() {
	var v TransactionsGetType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTransactionsGetType) SetTo(v TransactionsGetType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTransactionsGetType) Get() (v TransactionsGetType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTransactionsGetType) Or(d TransactionsGetType) TransactionsGetType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTransactionsTransactionIDPutReq returns new OptTransactionsTransactionIDPutReq with value set to v.
func NewOptTransactionsTransactionIDPutReq(v TransactionsTransactionIDPutReq) OptTransactionsTransactionIDPutReq {
	return OptTransactionsTransactionIDPutReq{
		Value: v,
		Set:   true,
	}
}

// OptTransactionsTransactionIDPutReq is optional TransactionsTransactionIDPutReq.
type OptTransactionsTransactionIDPutReq struct {
	Value TransactionsTransactionIDPutReq
	Set   bool
}

// IsSet returns true if OptTransactionsTransactionIDPutReq was set.
func (o OptTransactionsTransactionIDPutReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTransactionsTransactionIDPutReq) Reset() {
	var v TransactionsTransactionIDPutReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTransactionsTransactionIDPutReq) SetTo(v TransactionsTransactionIDPutReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTransactionsTransactionIDPutReq) Get() (v TransactionsTransactionIDPutReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTransactionsTransactionIDPutReq) Or(d TransactionsTransactionIDPutReq) TransactionsTransactionIDPutReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Transaction Extended Model.
// Ref: #/components/schemas/Transaction
type Transaction struct {
	ID        string           `json:"id"`
	State     TransactionState `json:"state"`
	From      TransactionFrom  `json:"from"`
	To        TransactionTo    `json:"to"`
	Amount    float64          `json:"amount"`
	Currency  Currency         `json:"currency"`
	CreatedAt string           `json:"created_at"`
}

// GetID returns the value of ID.
func (s *Transaction) GetID() string {
	return s.ID
}

// GetState returns the value of State.
func (s *Transaction) GetState() TransactionState {
	return s.State
}

// GetFrom returns the value of From.
func (s *Transaction) GetFrom() TransactionFrom {
	return s.From
}

// GetTo returns the value of To.
func (s *Transaction) GetTo() TransactionTo {
	return s.To
}

// GetAmount returns the value of Amount.
func (s *Transaction) GetAmount() float64 {
	return s.Amount
}

// GetCurrency returns the value of Currency.
func (s *Transaction) GetCurrency() Currency {
	return s.Currency
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Transaction) GetCreatedAt() string {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *Transaction) SetID(val string) {
	s.ID = val
}

// SetState sets the value of State.
func (s *Transaction) SetState(val TransactionState) {
	s.State = val
}

// SetFrom sets the value of From.
func (s *Transaction) SetFrom(val TransactionFrom) {
	s.From = val
}

// SetTo sets the value of To.
func (s *Transaction) SetTo(val TransactionTo) {
	s.To = val
}

// SetAmount sets the value of Amount.
func (s *Transaction) SetAmount(val float64) {
	s.Amount = val
}

// SetCurrency sets the value of Currency.
func (s *Transaction) SetCurrency(val Currency) {
	s.Currency = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Transaction) SetCreatedAt(val string) {
	s.CreatedAt = val
}

type TransactionFrom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *TransactionFrom) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *TransactionFrom) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *TransactionFrom) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *TransactionFrom) SetName(val string) {
	s.Name = val
}

// Transaction state.
// Ref: #/components/schemas/TransactionState
type TransactionState string

const (
	TransactionStatePENDING   TransactionState = "PENDING"
	TransactionStateCOMPLETED TransactionState = "COMPLETED"
	TransactionStateREJECTED  TransactionState = "REJECTED"
)

// AllValues returns all TransactionState values.
func (TransactionState) AllValues() []TransactionState {
	return []TransactionState{
		TransactionStatePENDING,
		TransactionStateCOMPLETED,
		TransactionStateREJECTED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TransactionState) MarshalText() ([]byte, error) {
	switch s {
	case TransactionStatePENDING:
		return []byte(s), nil
	case TransactionStateCOMPLETED:
		return []byte(s), nil
	case TransactionStateREJECTED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TransactionState) UnmarshalText(data []byte) error {
	switch TransactionState(data) {
	case TransactionStatePENDING:
		*s = TransactionStatePENDING
		return nil
	case TransactionStateCOMPLETED:
		*s = TransactionStateCOMPLETED
		return nil
	case TransactionStateREJECTED:
		*s = TransactionStateREJECTED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type TransactionTo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *TransactionTo) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *TransactionTo) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *TransactionTo) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *TransactionTo) SetName(val string) {
	s.Name = val
}

type TransactionsGetOK struct {
	Data  []Transaction `json:"data"`
	Count float64       `json:"count"`
	Skip  float64       `json:"skip"`
	Limit float64       `json:"limit"`
}

// GetData returns the value of Data.
func (s *TransactionsGetOK) GetData() []Transaction {
	return s.Data
}

// GetCount returns the value of Count.
func (s *TransactionsGetOK) GetCount() float64 {
	return s.Count
}

// GetSkip returns the value of Skip.
func (s *TransactionsGetOK) GetSkip() float64 {
	return s.Skip
}

// GetLimit returns the value of Limit.
func (s *TransactionsGetOK) GetLimit() float64 {
	return s.Limit
}

// SetData sets the value of Data.
func (s *TransactionsGetOK) SetData(val []Transaction) {
	s.Data = val
}

// SetCount sets the value of Count.
func (s *TransactionsGetOK) SetCount(val float64) {
	s.Count = val
}

// SetSkip sets the value of Skip.
func (s *TransactionsGetOK) SetSkip(val float64) {
	s.Skip = val
}

// SetLimit sets the value of Limit.
func (s *TransactionsGetOK) SetLimit(val float64) {
	s.Limit = val
}

type TransactionsGetOrderBy string

const (
	TransactionsGetOrderByAmount    TransactionsGetOrderBy = "amount"
	TransactionsGetOrderByCreatedAt TransactionsGetOrderBy = "created_at"
)

// AllValues returns all TransactionsGetOrderBy values.
func (TransactionsGetOrderBy) AllValues() []TransactionsGetOrderBy {
	return []TransactionsGetOrderBy{
		TransactionsGetOrderByAmount,
		TransactionsGetOrderByCreatedAt,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TransactionsGetOrderBy) MarshalText() ([]byte, error) {
	switch s {
	case TransactionsGetOrderByAmount:
		return []byte(s), nil
	case TransactionsGetOrderByCreatedAt:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TransactionsGetOrderBy) UnmarshalText(data []byte) error {
	switch TransactionsGetOrderBy(data) {
	case TransactionsGetOrderByAmount:
		*s = TransactionsGetOrderByAmount
		return nil
	case TransactionsGetOrderByCreatedAt:
		*s = TransactionsGetOrderByCreatedAt
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type TransactionsGetOrderDir string

const (
	TransactionsGetOrderDirAsc  TransactionsGetOrderDir = "asc"
	TransactionsGetOrderDirDesc TransactionsGetOrderDir = "desc"
)

// AllValues returns all TransactionsGetOrderDir values.
func (TransactionsGetOrderDir) AllValues() []TransactionsGetOrderDir {
	return []TransactionsGetOrderDir{
		TransactionsGetOrderDirAsc,
		TransactionsGetOrderDirDesc,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TransactionsGetOrderDir) MarshalText() ([]byte, error) {
	switch s {
	case TransactionsGetOrderDirAsc:
		return []byte(s), nil
	case TransactionsGetOrderDirDesc:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TransactionsGetOrderDir) UnmarshalText(data []byte) error {
	switch TransactionsGetOrderDir(data) {
	case TransactionsGetOrderDirAsc:
		*s = TransactionsGetOrderDirAsc
		return nil
	case TransactionsGetOrderDirDesc:
		*s = TransactionsGetOrderDirDesc
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type TransactionsGetType string

const (
	TransactionsGetTypeINGRESS TransactionsGetType = "INGRESS"
	TransactionsGetTypeEGRESS  TransactionsGetType = "EGRESS"
	TransactionsGetTypeALL     TransactionsGetType = "ALL"
)

// AllValues returns all TransactionsGetType values.
func (TransactionsGetType) AllValues() []TransactionsGetType {
	return []TransactionsGetType{
		TransactionsGetTypeINGRESS,
		TransactionsGetTypeEGRESS,
		TransactionsGetTypeALL,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TransactionsGetType) MarshalText() ([]byte, error) {
	switch s {
	case TransactionsGetTypeINGRESS:
		return []byte(s), nil
	case TransactionsGetTypeEGRESS:
		return []byte(s), nil
	case TransactionsGetTypeALL:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TransactionsGetType) UnmarshalText(data []byte) error {
	switch TransactionsGetType(data) {
	case TransactionsGetTypeINGRESS:
		*s = TransactionsGetTypeINGRESS
		return nil
	case TransactionsGetTypeEGRESS:
		*s = TransactionsGetTypeEGRESS
		return nil
	case TransactionsGetTypeALL:
		*s = TransactionsGetTypeALL
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type TransactionsTransactionIDPutCreated struct {
	Ok bool `json:"ok"`
}

// GetOk returns the value of Ok.
func (s *TransactionsTransactionIDPutCreated) GetOk() bool {
	return s.Ok
}

// SetOk sets the value of Ok.
func (s *TransactionsTransactionIDPutCreated) SetOk(val bool) {
	s.Ok = val
}

type TransactionsTransactionIDPutReq struct {
	FromAccountID string   `json:"from_account_id"`
	ToAccountID   string   `json:"to_account_id"`
	Amount        float64  `json:"amount"`
	Currency      Currency `json:"currency"`
}

// GetFromAccountID returns the value of FromAccountID.
func (s *TransactionsTransactionIDPutReq) GetFromAccountID() string {
	return s.FromAccountID
}

// GetToAccountID returns the value of ToAccountID.
func (s *TransactionsTransactionIDPutReq) GetToAccountID() string {
	return s.ToAccountID
}

// GetAmount returns the value of Amount.
func (s *TransactionsTransactionIDPutReq) GetAmount() float64 {
	return s.Amount
}

// GetCurrency returns the value of Currency.
func (s *TransactionsTransactionIDPutReq) GetCurrency() Currency {
	return s.Currency
}

// SetFromAccountID sets the value of FromAccountID.
func (s *TransactionsTransactionIDPutReq) SetFromAccountID(val string) {
	s.FromAccountID = val
}

// SetToAccountID sets the value of ToAccountID.
func (s *TransactionsTransactionIDPutReq) SetToAccountID(val string) {
	s.ToAccountID = val
}

// SetAmount sets the value of Amount.
func (s *TransactionsTransactionIDPutReq) SetAmount(val float64) {
	s.Amount = val
}

// SetCurrency sets the value of Currency.
func (s *TransactionsTransactionIDPutReq) SetCurrency(val Currency) {
	s.Currency = val
}
