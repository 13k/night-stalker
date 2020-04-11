package models

type OnBeforeCreate interface {
	BeforeCreate() error
}

type OnAfterCreate interface {
	AfterCreate() error
}

type OnBeforeUpdate interface {
	BeforeUpdate() error
}

type OnAfterUpdate interface {
	AfterUpdate() error
}

type OnBeforeSave interface {
	BeforeSave() error
}

type OnAfterSave interface {
	AfterSave() error
}

func BeforeCreate(r Record) error {
	if cb, ok := r.(OnBeforeSave); ok {
		if err := cb.BeforeSave(); err != nil {
			return err
		}
	}

	if cb, ok := r.(OnBeforeCreate); ok {
		if err := cb.BeforeCreate(); err != nil {
			return err
		}
	}

	return nil
}

func AfterCreate(r Record) error {
	if cb, ok := r.(OnAfterCreate); ok {
		if err := cb.AfterCreate(); err != nil {
			return err
		}
	}

	if cb, ok := r.(OnAfterSave); ok {
		if err := cb.AfterSave(); err != nil {
			return err
		}
	}

	return nil
}

func BeforeUpdate(r Record) error {
	if cb, ok := r.(OnBeforeSave); ok {
		if err := cb.BeforeSave(); err != nil {
			return err
		}
	}

	if cb, ok := r.(OnBeforeUpdate); ok {
		if err := cb.BeforeUpdate(); err != nil {
			return err
		}
	}

	return nil
}

func AfterUpdate(r Record) error {
	if cb, ok := r.(OnAfterUpdate); ok {
		if err := cb.AfterUpdate(); err != nil {
			return err
		}
	}

	if cb, ok := r.(OnAfterSave); ok {
		if err := cb.AfterSave(); err != nil {
			return err
		}
	}

	return nil
}
