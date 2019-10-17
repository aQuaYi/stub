# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 2.0.0 - 2019-10-16

### Added

- 添加了 `Var`，`Func`，`Env` 函数。
- 添加了 `Stubber` 接口

### Changed

把以前 `Stubs` 结构体的公开方法，移入了 `Stubber` 接口，并对名称进行了修改。

- `Stub` -> `Var`
- `StubFunc` -> `Func`
- `StubEnv` -> `Env`
- `Reset` -> `Restore`

## 1.0.0 - 2018-11-09

### Added

- Initial release with APIs to stub and reset values for testing.
- Supports stubbing Go variables, and environment variables.
