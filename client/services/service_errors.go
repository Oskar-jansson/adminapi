package services

import (
	"errors"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

// http error
var ErrClientNilResponse error = errors.New("nil response")
var ErrClientUnkownError error = errors.New("unkown error")
var ErrClientTrailingData error = errors.New("unexpected trailing data")
var ErrClientDecodeError error = errors.New("decode error")

// Api errors
var ErrInvalidUsernameOrPassword error = errors.New("1002")
var ErrAccesstokenInvalid error = errors.New("1003")
var ErrQueryStringInvalid error = errors.New("2001")
var ErrInputInvalid error = errors.New("2002")
var ErrParentReferanceInvalid error = errors.New("2003")
var ErrObjectIdExist error = errors.New("2004")
var ErrObjectIdDoesNotExist error = errors.New("2005")
var ErrRastampInvalid error = errors.New("2006")
var ErrObjectLinked error = errors.New("2007")
var ErrMissingId error = errors.New("2008")
var ErrDatabaseError error = errors.New("3001")
var ErrNotImplimented error = errors.New("3002")
var ErrOtherError error = errors.New("4000")
var ErrTimestampInvalid error = errors.New("8001")
var ErrTranslationConfigurationInvalid error = errors.New("8002")
var ErrInvalidConfigurationFile error = errors.New("8003")
var ErrValidationFalied error = errors.New("8004")
var ErrControllerUnitOffline error = errors.New("8005")
var ErrNoFreeOperatorLicens error = errors.New("8006")
var ErrAccessDenied error = errors.New("8007")
var ErrMissingApiLicens error = errors.New("8008")
var ErrSystemInvalid error = errors.New("8009")
var ErrServiceOffline error = errors.New("8010")
var ErrInvalidApiKeyOrLoginsExceded error = errors.New("8011")

func parseApiErrorCode(errMsg *models.ErrorMessage) error {

	err := fmt.Errorf("%s", errMsg.Friendlymessage)

	if errMsg == nil {
		return errors.New("error message was nil")
	}

	switch errMsg.Errorcode {
	case 1002:
		return errors.Join(err, ErrInvalidUsernameOrPassword)
	case 1003:
		return errors.Join(err, ErrAccesstokenInvalid)
	case 2001:
		return errors.Join(err, ErrQueryStringInvalid)
	case 2002:
		return errors.Join(err, ErrInputInvalid)
	case 2003:
		return errors.Join(err, ErrParentReferanceInvalid)
	case 2004:
		return errors.Join(err, ErrObjectIdExist)
	case 2005:
		return errors.Join(err, ErrObjectIdDoesNotExist)
	case 2006:
		return errors.Join(err, ErrRastampInvalid)
	case 2007:
		return errors.Join(err, ErrObjectLinked)
	case 2008:
		return errors.Join(err, ErrMissingId)
	case 3001:
		return errors.Join(err, ErrDatabaseError)
	case 3002:
		return errors.Join(err, ErrNotImplimented)
	case 4000:
		return errors.Join(err, ErrOtherError)
	case 8001:
		return errors.Join(err, ErrTimestampInvalid)
	case 8002:
		return errors.Join(err, ErrTranslationConfigurationInvalid)
	case 8003:
		return errors.Join(err, ErrInvalidConfigurationFile)
	case 8004:
		return errors.Join(err, ErrValidationFalied)
	case 8005:
		return errors.Join(err, ErrControllerUnitOffline)
	case 8006:
		return errors.Join(err, ErrNoFreeOperatorLicens)
	case 8007:
		return errors.Join(err, ErrAccessDenied)
	case 8008:
		return errors.Join(err, ErrMissingApiLicens)
	case 8009:
		return errors.Join(err, ErrSystemInvalid)
	case 8010:
		return errors.Join(err, ErrServiceOffline)
	case 8011:
		return errors.Join(err, ErrInvalidApiKeyOrLoginsExceded)
	default:
		return ErrOtherError
	}
}

// Reject codes
var ErrHasReject error = errors.New("rejected") // generic indication of rejection
var ErrRequiredDataMissingInField error = errors.New("1")
var ErrStringOrNumberToLarge error = errors.New("2")
var ErrStringOrNumberToSmall error = errors.New("3")
var ErrInvalidCharacters error = errors.New("4")
var ErrObjectOrUniqeIdentifierAlreadyExists error = errors.New("5")
var ErrInvalidCombinationWithSystemPolicy error = errors.New("6")

func parseApiRejectCode(errMsg *models.ErrorMessage) error {

	var rejectErrors error

	for _, reject := range *errMsg.Reject {
		switch reject.Code {
		case 1:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' is missing required data.", reject.Name))
		case 2:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' contains a value that is too large.", reject.Name))
		case 3:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' contains a value that is too small.", reject.Name))
		case 4:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' contains invalid characters.", reject.Name))
		case 5:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' or its unique identifier already exists.", reject.Name))
		case 6:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' has an invalid combination with system policy.", reject.Name))
		default:
			rejectErrors = errors.Join(rejectErrors, fmt.Errorf("Field '%s' has an unknown reject code: %d.", reject.Name, reject.Code))
		}
	}

	if rejectErrors != nil {
		rejectErrors = errors.Join(rejectErrors, ErrHasReject)
	}

	return rejectErrors
}

// parses whole error message and returns a chain of errors
func parseErrorMessage(errMsg *models.ErrorMessage) error {
	return errors.Join(parseApiErrorCode(errMsg), parseApiRejectCode(errMsg))
}
