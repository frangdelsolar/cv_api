package displayData_service

import (
	m "cv_api/models"
	ddRepository "cv_api/repositories/displayData.repository"
)

func Create(dd m.DisplayData) error {

	err := ddRepository.Create(dd)

	if err != nil {
		return err
	}

	return nil
}

func Read(userId string) (m.DisplayData, error) {

	dd, err := ddRepository.Read(userId)

	if err != nil {
		return dd, err
	}

	return dd, nil
}

// func Update(user m.DisplayData, userId string) error {

// 	err := ddRepository.Update(user, userId)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func Delete(userId string) error {

// 	err := ddRepository.Delete(userId)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
