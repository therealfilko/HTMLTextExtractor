package mock

import (
    "errors"
)

type MockStorage struct {
    Data      string
    ShouldFail bool
}

func (m *MockStorage) Save(data string) error {
    if m.ShouldFail {
        return errors.New("mock error")
    }
    m.Data = data
    return nil
}
