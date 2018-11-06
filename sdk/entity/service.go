package entity

import "context"

const (
    CREATE = "create"
    GET_BY_ID = "get_by_id"
    SET_STATUS = "set_status"
    SET_NAME = "set_name"
    SET_URL = "set_url"
    ADD_CHILD_BY_ID = "add_child_by_id"

    OP_DONE = "op_done"
)

type Service struct {
    storages *StorageStack
    continueOnError bool
}

func NewService(storage StorageInterface, coe bool) *Service {
    return &Service{
        storages: NewStorageStack(storage),
        continueOnError: coe,
    }
}

func (s *Service) AppendStorage(storage StorageInterface) {
    s.storages.Append(storage)
}

func (s *Service) GetStorages() *StorageStack {
    return s.storages
}

func (s *Service) Create(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, CREATE, input, output)
}

func (s *Service) SetStatus(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, SET_STATUS, input, output)
}

func (s *Service) SetName(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, SET_NAME, input, output)
}

func (s *Service) SetUrl(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, SET_URL, input, output)
}

func (s *Service) AddChildById(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, ADD_CHILD_BY_ID, input, output)
}

func (s *Service) GetById(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, GET_BY_ID, input, output)
}

func (s *Service) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
    defer func() {
        go s.afterOp(ctx, opName, input, output)
    }()

    opDone := false
    doneCtx := context.WithValue(ctx, OP_DONE, &opDone)

    st := s.GetStorages()
    for st != nil {
        err = st.GetStorage().Op(doneCtx, opName, input, output)
        if err != nil && !s.continueOnError {
            return
        }

        if opDone {
            return
        }

        st = st.Next
    }

    return
}

func (s *Service) afterOp(ctx context.Context, opName string, input interface{}, output interface{}) {
    opDone := false
    doneCtx := context.WithValue(ctx, OP_DONE, &opDone)

    st := s.GetStorages()
    for st != nil {
        st.GetStorage().AfterOp(doneCtx, opName, input, output)
        if opDone {
            return
        }

        st = st.Next
    }

    return
}