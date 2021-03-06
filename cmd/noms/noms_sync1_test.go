// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"path"
	"testing"

	"github.com/attic-labs/noms/go/chunks"
	"github.com/attic-labs/noms/go/d"
	"github.com/attic-labs/noms/go/datas"
	"github.com/attic-labs/noms/go/dataset"
	"github.com/attic-labs/noms/go/spec"
	"github.com/attic-labs/noms/go/types"
	"github.com/attic-labs/noms/go/util/clienttest"
	"github.com/attic-labs/testify/suite"
)

func TestAgSync(t *testing.T) {
	d.UtilExiter = testExiter{}
	suite.Run(t, &nomsAgSyncTestSuite{})
}

type nomsAgSyncTestSuite struct {
	clienttest.ClientTestSuite
}

func (s *nomsAgSyncTestSuite) TestSync() {
	source1 := dataset.NewDataset(datas.NewDatabase(chunks.NewLevelDBStore(s.LdbDir, "", 1, false)), "src")
	log.Printf("Dataset ID = %s",source1.ID())
	source1, err := source1.CommitValue(types.Number(42))
	s.NoError(err)
	source2, err := source1.CommitValue(types.Number(43))
	s.NoError(err)
	source1HeadRef := source1.Head().Hash()
	source2.Database().Close() // Close Database backing both Datasets

	sourceSpec := spec.CreateValueSpecString("ldb", s.LdbDir, "#"+source1HeadRef.String())
	ldb2dir := path.Join(s.TempDir, "ldb2")
	sinkDatasetSpec := spec.CreateValueSpecString("ldb", ldb2dir, "dest")
	sout, _ := s.Run(main, []string{"sync", sourceSpec, sinkDatasetSpec})
	log.Print(sout)

	s.Regexp("Created", sout)
	dest := dataset.NewDataset(datas.NewDatabase(chunks.NewLevelDBStore(ldb2dir, "", 1, false)), "dest")
	s.True(types.Number(42).Equals(dest.HeadValue()))
	dest.Database().Close()

	sourceDataset := spec.CreateValueSpecString("ldb", s.LdbDir, "src")
	sout, _ = s.Run(main, []string{"sync", sourceDataset, sinkDatasetSpec})
	s.Regexp("Synced", sout)

	dest = dataset.NewDataset(datas.NewDatabase(chunks.NewLevelDBStore(ldb2dir, "", 1, false)), "dest")
	s.True(types.Number(43).Equals(dest.HeadValue()))
	dest.Database().Close()

	sout, _ = s.Run(main, []string{"sync", sourceDataset, sinkDatasetSpec})
	s.Regexp("up to date", sout)
}
