package registry

import (
	"fmt"
	"github.com/layer5io/meshkit/errors"
)

var (
	ErrAppendToSheetCode      = "meshkit-11261"
	ErrGenerateModelCode      = "meshkit-11262"
	ErrMarshalStructToCSVCode = "meshkit-11263"
)

func ErrGenerateModel(err error, modelName string) error {
	return errors.New(ErrGenerateModelCode, errors.Alert, []string{fmt.Sprintf("error generating model: %s", modelName)}, []string{fmt.Sprintf("Error generating model: %s\n %s", modelName, err.Error())}, []string{"Registrant used for the model is not supported", "Verify the model's source URL.", "Failed to create a local directory in the filesystem for this model."}, []string{"Ensure that each kind of registrant used is a supported kind.", "Ensure correct model source URL is provided and properly formatted.", "Ensure sufficient permissions to allow creation of model directory."})
}

func ErrAppendToSheet(err error, id string) error {
	return errors.New(ErrAppendToSheetCode, errors.Alert,
		[]string{fmt.Sprintf("Failed to append data into sheet %s", id)},
		[]string{err.Error()},
		[]string{"Error occurred while appending to the spreadsheet", "The credential might be incorrect/expired"},
		[]string{"Ensure correct append range (A1 notation) is used", "Ensure correct credential is used"})
}

func ErrMarshalStructToCSV(err error) error {
	return errors.New(ErrMarshalStructToCSVCode, errors.Alert,
		[]string{"Failed to marshal struct to csv"},
		[]string{err.Error()},
		[]string{"The column names in your spreadsheet do not match the names in the struct.", " For example, the spreadsheet has a column named 'First Name' but the struct expects a column named 'firstname'. Please make sure the names match exactly."},
		[]string{"The column names in the spreadsheet do not match the names in the struct. Please make sure they are spelled exactly the same and use the same case (uppercase/lowercase).", "The value you are trying to convert is not of the expected type for the column. Please ensure it is a [number, string, date, etc.].", "The column names in your spreadsheet do not match the names in the struct. For example, the spreadsheet has a column named 'First Name' but the struct expects a column named 'firstname'. Please make sure the names match exactly."})
}
