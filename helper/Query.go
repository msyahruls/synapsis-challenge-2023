package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category
func QueryCategory(issueId string, name string) primitive.M {
	query := bson.M{}

	if issueId != "" {
		query["_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
	}

	if name != "" {
		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}

	return query
}

// Category Exist
func QueryCategoryExist(id string, name string) primitive.M {
	query := bson.M{
		"$or": []interface{}{
			bson.M{"_id": id},
			bson.M{"name": name},
		},
	}

	return query
}

// Product
func QueryProduct(issueId string, name string, category string) primitive.M {
	query := bson.M{}

	if issueId != "" {
		query["_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
	}

	if name != "" {
		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}

	if category != "" {
		query["category_id"] = primitive.Regex{Pattern: category, Options: "i"}
	}

	return query
}

// Product Exist
func QueryProductExist(name string) primitive.M {
	query := bson.M{
		"$or": []interface{}{
			bson.M{"name": name},
		},
	}

	return query
}

// Cart
func QueryCart(UserID string, ProductID string) primitive.M {
	query := bson.M{}

	if UserID != "" {
		query["user_id"] = primitive.Regex{Pattern: UserID, Options: "i"}
	}

	if ProductID != "" {
		query["product_id"] = primitive.Regex{Pattern: ProductID, Options: "i"}
	}

	return query
}

// Cart Exist
func QueryCartExist(UserID string, ProductID string) primitive.M {
	query := bson.M{
		"$and": []interface{}{
			bson.M{"user_id": UserID},
			bson.M{"prodcut_id": ProductID},
		},
	}

	return query
}

// Transaction
func QueryTransaction(UserID string) primitive.M {
	query := bson.M{}

	if UserID != "" {
		query["user_id"] = primitive.Regex{Pattern: UserID, Options: "i"}
	}

	return query
}

// Transaction Exist
func QueryTransactionExist(UserID string) primitive.M {
	query := bson.M{
		"$and": []interface{}{
			bson.M{"user_id": UserID},
		},
	}

	return query
}
