<center>
  <h1>Rhyolite</h1>
</center>

🚧 **Rhyolite is under heavy development and is not currently ready for end-users.**

## Usage
Rhyolite requires an initialised sqlite database. This can be done using
```sh
rhyolite init
```
<sup>This applies the schema provided in `./database/schema.sql` as pointed to by `./sqlc.yml`.</sup>

Once the database is initialised, notes can be created using the following command
```sh
rhyolite create "Note title" "Note content"
```

and can be listed using
```
rhyolite list
```


## Development Roadmap
- [x] Database creation
- [ ] CRUD
  - [x] Create Notes
  - [x] Read Notes
  - [ ] Update Notes
  - [ ] Delete Notes
- [ ] Timestamps
- [ ] TUI Interface (bubbletea)
- [ ] Additional, extra features
