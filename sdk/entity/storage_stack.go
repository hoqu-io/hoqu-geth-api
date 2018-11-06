package entity

type StorageStack struct {
    Next *StorageStack
    storage StorageInterface
}

func NewStorageStack(storage StorageInterface) *StorageStack {
    return &StorageStack{
        Next: nil,
        storage: storage,
    }
}

func (s *StorageStack) Append(storage StorageInterface) {
    st := &StorageStack{
        Next: nil,
        storage: storage,
    }
    s.Tail().Next = st
}

func (s *StorageStack) Tail() *StorageStack {
    t := s
    for t.Next != nil {
        t = t.Next
    }

    return t
}

func (s *StorageStack) GetStorage() StorageInterface {
    return s.storage
}