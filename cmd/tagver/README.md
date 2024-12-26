# tagver
Tagver manages the boilerplate of managing semantic versions in git tags.

## Usage
> [!IMPORTANT]
> tagver does NOT make changes. It only gets, or previews the semantic version of tags of git repositories.

Assume we are in a git repository which is at version tag `v2.1.4`.

Get the current version:
```sh
tagver # v2.1.4
```
Get what the next patch version would be:
```sh
tagver patch # v2.1.5
```
Get what the next minor version would be:
```sh
agver minor # v2.2.0
```
Get what the next major version would be:
```sh
tagver major # v3.0.0
```
Usage with git:
```sh
# make a new tag with the next patch version.
git tag (tagver patch) # v2.1.4
# push the current, v2.1.4, tag to origin.
git push origin (tagver) 
```
## License
tagver is licensed under GPL 3.0.
