package userapplication

func AddTravel(travelID *string, email *string) error {
	err := Repo.AddTravel(travelID, email)
	if err != nil {
		return err
	}

	return nil
}
