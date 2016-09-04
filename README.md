# Noms Extended Test Suite Files

These are additional tests that you may want to run with
[Noms]
(https://github.com/attic-labs/noms)

Start off by cloning this repo if you plan on making changes to it.
Copy the test files over to the noms source tree by running this command.
This will pollute your noms src tree with these files, there is a script
to remove these files once you have ran the extra tests.

```
./acp
```

```
gtv -run 'Ag'
```

If you make any modifications to your test files you can copy them
back over to your repo and then commit them back to github.

```
./acpback
```

And then if you no longer want the extra tests in the noms repo you can
go ahead and delete them.

```
./adel
```
