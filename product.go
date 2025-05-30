// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import "encoding/json"

type Product struct {
	Id                                          string         `json:"id"`
	Code                                        string         `json:"code"`
	Brands                                      string         `json:"brands"`
	BrandsTags                                  []string       `json:"brands_tags"`
	GenericName                                 string         `json:"generic_name"`
	GenericNameEn                               string         `json:"generic_name_en"`
	CreatedTime                                 EpochTime      `json:"created_t"`
	LastImageTime                               EpochTime      `json:"last_image_t"`
	LastModifiedTime                            EpochTime      `json:"last_modified_t"`
	CompletedTime                               EpochTime      `json:"completed_t"`
	ImageFrontSmallURL                          URL            `json:"image_front_smallURL"`
	ImageFrontThumbURL                          URL            `json:"image_front_thumb_url"`
	ImageFrontURL                               URL            `json:"image_front_url"`
	ImageIngredientsSmallURL                    URL            `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL                    URL            `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL                         URL            `json:"image_ingredients_url"`
	ImageNutritionSmallURL                      URL            `json:"image_nutrition_small_url"`
	ImageNutritionThumbURL                      URL            `json:"image_nutrition_thumb_url"`
	ImageNutritionURL                           URL            `json:"image_nutrition_url"`
	ImageSmallURL                               URL            `json:"image_small_url"`
	ImageThumbURL                               URL            `json:"image_thumb_url"`
	ImageURL                                    URL            `json:"image_url"`
	Creator                                     string         `json:"creator"`
	Checkers                                    []interface{}  `json:"checkers"`
	CheckersTags                                []interface{}  `json:"checkers_tags"`
	Editors                                     []string       `json:"editors"`
	EditorsTags                                 []string       `json:"editors_tags"`
	Correctors                                  []string       `json:"correctors"`
	CorrectorsTags                              []string       `json:"correctors_tags"`
	Informers                                   []string       `json:"informers"`
	InformersTags                               []string       `json:"informers_tags"`
	Photographers                               []string       `json:"photographers"`
	PhotographersTags                           []string       `json:"photographers_tags"`
	LastEditDatesTags                           []string       `json:"last_edit_dates_tags"`
	LastEditor                                  string         `json:"last_editor"`
	LastImageDatesTags                          []string       `json:"last_image_dates_tags"`
	LastModifiedBy                              string         `json:"last_modified_by"`
	Additives                                   string         `json:"additives"`
	AdditivesDebugTags                          []interface{}  `json:"additives_debug_tags"`
	AdditivesNumber                             int            `json:"additives_n"`
	AdditivesOldNumber                          int            `json:"additives_old_n"`
	AdditivesOldTags                            []string       `json:"additives_old_tags"`
	AdditivesPrev                               string         `json:"additives_prev"`
	AdditivesPrevNumber                         int            `json:"additives_prev_n"`
	AdditivesPrevTags                           []string       `json:"additives_prev_tags"`
	AdditivesTags                               []string       `json:"additives_tags"`
	AdditivesTagsNumber                         interface{}    `json:"additives_tags_n"`
	Allergens                                   string         `json:"allergens"`
	AllergensHierarchy                          []interface{}  `json:"allergens_hierarchy"`
	AllergensTags                               []interface{}  `json:"allergens_tags"`
	Categories                                  string         `json:"categories"`
	CategoriesDebugTags                         []interface{}  `json:"categories_debug_tags"`
	CategoriesHierarchy                         []string       `json:"categories_hierarchy"`
	CategoriesPrevHierarchy                     []string       `json:"categories_prev_hierarchy"`
	CategoriesPrevTags                          []string       `json:"categories_prev_tags"`
	CategoriesTags                              []string       `json:"categories_tags"`
	CitiesTags                                  []interface{}  `json:"cities_tags"`
	CodesTags                                   []string       `json:"codes_tags"`
	Complete                                    int            `json:"complete"`
	Countries                                   string         `json:"countries"`
	CountriesHierarchy                          []string       `json:"countries_hierarchy"`
	CountriesTags                               []string       `json:"countries_tags"`
	EmbCodes                                    string         `json:"emb_codes"`
	EmbCodes20141016                            string         `json:"emb_codes_20141016"`
	EmbCodesOrig                                string         `json:"emb_codes_orig"`
	EmbCodesTags                                []interface{}  `json:"emb_codes_tags"`
	EntryDatesTags                              []string       `json:"entry_dates_tags"`
	EcoscoreGrade                               string         `json:"ecoscore_grade"`
	EcoscoreTags                                []string       `json:"ecoscore_tags"`
	ExpirationDate                              string         `json:"expiration_date"`
	FruitsVegetablesNuts100GEstimate            json.Number    `json:"fruits-vegetables-nuts_100g_estimate"`
	Ingredients                                 []Ingredient   `json:"ingredients"`
	IngredientsDebug                            []interface{}  `json:"ingredients_debug"`
	IngredientsFromOrThatMayBeFromPalmOilNumber int            `json:"ingredients_from_or_that_may_be_from_palm_oil_n"`
	IngredientsFromPalmOilNumber                int            `json:"ingredients_from_palm_oil_n"`
	IngredientsFromPalmOilTags                  []interface{}  `json:"ingredients_from_palm_oil_tags"`
	IngredientsIdsDebug                         []string       `json:"ingredients_ids_debug"`
	IngredientsN                                json.Number    `json:"ingredients_n"`
	IngredientsNTags                            []string       `json:"ingredients_n_tags"`
	IngredientsTags                             []string       `json:"ingredients_tags"`
	IngredientsText                             string         `json:"ingredients_text"`
	IngredientsTextDebug                        string         `json:"ingredients_text_debug"`
	IngredientsTextEn                           string         `json:"ingredients_text_en"`
	IngredientsTextWithAllergens                string         `json:"ingredients_text_with_allergens"`
	IngredientsTextWithAllergensEn              string         `json:"ingredients_text_with_allergens_en"`
	IngredientsThatMayBeFromPalmOilNumber       int            `json:"ingredients_that_may_be_from_palm_oil_n"`
	IngredientsThatMayBeFromPalmOilTags         []interface{}  `json:"ingredients_that_may_be_from_palm_oil_tags"`
	InterfaceVersionCreated                     string         `json:"interface_version_created"`
	InterfaceVersionModified                    string         `json:"interface_version_modified"`
	Keywords                                    []string       `json:"_keywords"`
	Labels                                      string         `json:"labels"`
	LabelsDebugTags                             []interface{}  `json:"labels_debug_tags"`
	LabelsHierarchy                             []string       `json:"labels_hierarchy"`
	LabelsPrevHierarchy                         []string       `json:"labels_prev_hierarchy"`
	LabelsPrevTags                              []string       `json:"labels_prev_tags"`
	LabelsTags                                  []string       `json:"labels_tags"`
	Lang                                        string         `json:"lang"`
	Languages                                   map[string]int `json:"languages"`
	LanguagesCodes                              map[string]int `json:"languages_codes"`
	LanguagesHierarchy                          []string       `json:"languages_hierarchy"`
	LanguagesTags                               []string       `json:"languages_tags"`
	Locale                                      string         `json:"lc"`
	ManufacturingPlaces                         string         `json:"manufacturing_places"`
	ManufacturingPlacesTags                     []string       `json:"manufacturing_places_tags"`
	MaxImgId                                    string         `json:"max_imgid"`
	NewAdditivesNumber                          int            `json:"new_additives_n"`
	NoNutritionData                             interface{}    `json:"no_nutrition_data"`
	NutrientLevels                              NutrientLevel  `json:"nutrient_levels"`
	NutrientLevelsTags                          []string       `json:"nutrient_levels_tags"`
	Nutriments                                  Nutriment      `json:"nutriments"`
	NutritionDataPer                            string         `json:"nutrition_data_per"`
	NutritionGradeFr                            string         `json:"nutrition_grade_fr"`
	NutritionGrades                             string         `json:"nutrition_grades"`
	NutritionGradesTags                         []string       `json:"nutrition_grades_tags"`
	NutritionScoreDebug                         string         `json:"nutrition_score_debug"`
	OriginTags                                  []string       `json:"origins_tags"`
	Origins                                     string         `json:"origins"`
	Packaging                                   string         `json:"packaging"`
	PackagingTags                               []string       `json:"packaging_tags"`
	PnnsGroups1                                 string         `json:"pnns_groups_1"`
	PnnsGroups1Tags                             []string       `json:"pnns_groups_1_tags"`
	PnnsGroups2                                 string         `json:"pnns_groups_2"`
	PnnsGroups2Tags                             []string       `json:"pnns_groups_2_tags"`
	ProductName                                 string         `json:"product_name"`
	ProductNameEn                               string         `json:"product_name_en"`
	ProductNameDe                               string         `json:"product_name_de"`
	PurchasePlaces                              string         `json:"purchase_places"`
	PurchasePlacesTags                          []interface{}  `json:"purchase_places_tags"`
	Quantity                                    string         `json:"quantity"`
	Rev                                         json.Number    `json:"rev"`
	ScansNumber                                 int            `json:"scans_n"`
	ServingQuantity                             json.Number    `json:"serving_quantity"`
	ServingSize                                 string         `json:"serving_size"`
	SortKey                                     int            `json:"sortkey"`
	States                                      string         `json:"states"`
	StatesHierarchy                             []string       `json:"states_hierarchy"`
	StatesTags                                  []string       `json:"states_tags"`
	Stores                                      string         `json:"stores"`
	StoresTags                                  []interface{}  `json:"stores_tags"`
	Traces                                      string         `json:"traces"`
	TracesHierarchy                             []string       `json:"traces_hierarchy"`
	TracesTags                                  []string       `json:"traces_tags"`
	UniqueScansNumber                           int            `json:"unique_scans_n"`
	UnknownNutrientsTags                        []interface{}  `json:"unknown_nutrients_tags"`
	UpdateKey                                   string         `json:"update_key"`
}
