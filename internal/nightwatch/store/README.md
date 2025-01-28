# Store 层

因为 Store 层代码相对不易变，所以，Store 层很少会随着项目迭代，衍生出 V2 版本，因此 Store 层只需要一个版本即可。

## Store 层开发规范

- 每一个数据库表对应一个 Store 层资源，资源的CURD代码按资源分文件保存；
- 每个 Store 层资源方法均包括标准的CURD方法和扩展方法，例如：

```go
// PostStore 定义了 post 模块在 store 层所实现的方法.
type PostStore interface {
    Create(ctx context.Context, obj *model.PostM) error
    Update(ctx context.Context, obj *model.PostM) error
    Delete(ctx context.Context, opts *where.Options) error
    Get(ctx context.Context, opts *where.Options) (*model.PostM, error)
    List(ctx context.Context, opts *where.Options) (int64, []*model.PostM, error)

    PostExpansion
}

// PostExpansion 定义了帖子操作的附加方法.
type PostExpansion interface{}
```

- Store 层方法继承自 genericstore，如果需要自定义逻辑，可重新方法。
