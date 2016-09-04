# Noms Extended Test Suite Files

These are additional tests that you may want to run with
[Noms]
(https://github.com/attic-labs/noms)

Start off by cloning this repo if you plan on making changes to it.

Copy the test files over to the noms source tree by running this command.

This will pollute your noms src tree with the extra test files, there is a script
called **adel** to remove these files once you have run the extra tests.

```
./acp
```

Now you can run only these test files by going to the directory
you are interested in testing further.

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
