# Commit Message Guidelines

## Commit Message Format

Each commit message consists of a **header**, a **body**, and one or more **footer**(s). The header has a special format that includes a **type**, a **scope**, and a **description**:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

The **header** is mandatory; the **scope** of the header is optional. The **body** and **footer**(s) are also optional.

Any line of the commit message cannot be longer than 80 characters. This allows for the message to be easier read in various git tools.

### Commit Types

- `fix`: a commit of the *type* `fix` patches a bug in your codebase (this correlates with `PATCH` in [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
- `feat`: a commit of the *type* `feat` introduces a new feature to the codebase (this correlates with `MINOR` in Semantic Versioning).

#### Additional Commit Types

- `build`: Changes that affect the build system or external dependencies (example scopes: make, gulp, grunt)
- `chore`: Changes that benefit miscellaneous tasks like updating internal dependencies (example scopes: go.mod, package.json, requirements.txt)
- `deploy`: Changes to our deployment configuration files and scripts (example scopes: Travis, Circle, Docker)
- `docs`: Documentation only changes (example scopes: README.md, CHANGELOG.md)
- `perf`: A code change that improves performance 
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `test`: Adding missing tests or correcting existing tests

Additional types are not mandated by the Conventional Commits specification, and have no implicit effect in Semantic Versioning (unless they include a `BREAKING CHANGE`). A scope may be provided to a commit’s type, to provide additional contextual information and is contained within parenthesis, e.g., `feat(parser): add ability to parse arrays`.

### Commit Footers

- `BREAKING-CHANGE`: also indicated by an `!` after the **scope** in the **header**; a commit that has a `BREAKING-CHANGE` footer introduces a change to the codebase that impacts client usage (this correlates with `MAJOR` in Semantic Versioning). A `BREAKING-CHANGE` can be part of commits of any *type*.
- *footers* other than `BREAKING-CHANGE: <description>` may be provided and follow a convention similar to [git trailer format](https://git-scm.com/docs/git-interpret-trailers).

## Examples

### Commit message with a description and a breaking change footer

```
feat: allow provided config object to extend other configs

BREAKING-CHANGE: `extends` key in config file is now used for extending other config files
```

### Commit message with `!` to draw attention to breaking change

```
feat!: send an email to the customer when a product is shipped
```

### Commit message with scope and `!` to draw attention to breaking change

```
feat(api)!: send an email to the customer when a product is shipped
```

### Commit message with both `!` and `BREAKING-CHANGE` footer

```
chore!: drop support for Node 6

BREAKING-CHANGE: use JavaScript features not available in Node 6.
```

### Commit message with no body

```
docs: correct spelling of CHANGELOG
```

### Commit message with scope

```
feat(lang): add Polish language
```

### Commit message with  multi-paragraph body and multiple footers

```
fix: prevent racing of requests

Introduce a request id and a reference to latest request. Dismiss
incoming responses other than from latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

Reviewed-by: Z
Refs: #123
```

The above examples are taken directly from the Conventional Commits [Examples](https://www.conventionalcommits.org/en/v1.0.0/#example) section. Consider looking through the Conventional Commits [Specification](https://www.conventionalcommits.org/en/v1.0.0/#specification) for a better understanding.

## Attribution

This guideline is based on both the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) style and the [Angular convention](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines). This leads to **more readable messages** that are easy to follow when looking through the **project history**.

