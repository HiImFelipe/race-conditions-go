## Avoiding race conditions in GO

Two examples on how to avoid race conditions in GoLang.
One of them uses the Atomic API, the other uses regular Mutex principles.

`go work` was used in order to have two projects in the same root folder.

### Running

Requirements
- GoLang 1.22 or higher;

```bash
go run atomic/main.go

go run mutex/main.go 
```

### References

- <a href="https://www.youtube.com/watch?v=egd4WHJMwC0">Cap. 18 – Concorrência – 5. Mutex</a>
- <a href="https://www.youtube.com/watch?v=iFlQ2yAYcp4">Cap. 18 – Concorrência – 6. Atomic</a>
