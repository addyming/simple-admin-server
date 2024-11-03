// Code generated by ent, DO NOT EDIT.

package emstockwatch

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the emstockwatch type in the database.
	Label = "em_stock_watch"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "stock_watch_id"
	// FieldStockWatchCode holds the string denoting the stock_watch_code field in the database.
	FieldStockWatchCode = "stock_watch_code"
	// FieldStockWatchName holds the string denoting the stock_watch_name field in the database.
	FieldStockWatchName = "stock_watch_name"
	// FieldStockWatchUpPrice holds the string denoting the stock_watch_up_price field in the database.
	FieldStockWatchUpPrice = "stock_watch_up_price"
	// FieldStockWatchFallPrice holds the string denoting the stock_watch_fall_price field in the database.
	FieldStockWatchFallPrice = "stock_watch_fall_price"
	// FieldStockWatchUpChange holds the string denoting the stock_watch_up_change field in the database.
	FieldStockWatchUpChange = "stock_watch_up_change"
	// FieldStockWatchFallChange holds the string denoting the stock_watch_fall_change field in the database.
	FieldStockWatchFallChange = "stock_watch_fall_change"
	// Table holds the table name of the emstockwatch in the database.
	Table = "em_stock_watch"
)

// Columns holds all SQL columns for emstockwatch fields.
var Columns = []string{
	FieldID,
	FieldStockWatchCode,
	FieldStockWatchName,
	FieldStockWatchUpPrice,
	FieldStockWatchFallPrice,
	FieldStockWatchUpChange,
	FieldStockWatchFallChange,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the EmStockWatch queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStockWatchCode orders the results by the stock_watch_code field.
func ByStockWatchCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchCode, opts...).ToFunc()
}

// ByStockWatchName orders the results by the stock_watch_name field.
func ByStockWatchName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchName, opts...).ToFunc()
}

// ByStockWatchUpPrice orders the results by the stock_watch_up_price field.
func ByStockWatchUpPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchUpPrice, opts...).ToFunc()
}

// ByStockWatchFallPrice orders the results by the stock_watch_fall_price field.
func ByStockWatchFallPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchFallPrice, opts...).ToFunc()
}

// ByStockWatchUpChange orders the results by the stock_watch_up_change field.
func ByStockWatchUpChange(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchUpChange, opts...).ToFunc()
}

// ByStockWatchFallChange orders the results by the stock_watch_fall_change field.
func ByStockWatchFallChange(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockWatchFallChange, opts...).ToFunc()
}
