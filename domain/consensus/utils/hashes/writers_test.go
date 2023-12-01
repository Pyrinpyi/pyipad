package hashes

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewBlockHash(t *testing.T) {
	datas := [][]byte{
		{},
		{1},
		{5, 199, 126, 44, 71, 32, 82, 139, 122, 217, 43, 48, 52, 112, 40, 209, 180, 83, 139, 231, 72, 48, 136, 48, 168, 226, 133, 7, 60, 4, 160, 205},
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	tests := []struct {
		hasher   HashWriter
		expected []string
	}{
		{NewTransactionHashWriter(), []string{
			"bdecbc3411e21e9b1c9f45d3def8a06e7c498a8b06c0300f8d4b4ec0d21dbd9c",
			"50ec906c2f9463559a13d7625f39f23ea2f395ef8923ffc3cf704ce326bcd79d",
			"cee07d0d436d9625edfb435fea908df006af474b71e18eefdd46403e7028679b",
			"a86c35fa0dd47963bcfaefed57198c4ef2b3c5db9c810a4d96e12d0ad9bf62a2",
			"4225a247753d9de9855bb4a15b510808de920957ab7ccd3b4fd4adff47e9567c",
		}},
		{NewTransactionIDWriter(), []string{
			"332f3ef084d316f53440623d784a15bd215b098bfcbaaef96ba06cf6ce1e6b5e",
			"80f5f4073f8cfd22978aa0cdb5da1c795b29d4ded81a26da2a9a4c1f54431487",
			"ba7cedb37f6c64053d94182f3e545c657c872191af709bcc6ce0be837ead10d6",
			"a6f72ce2928cc0603eeabad8931cc4aab050c28a9e120e5d0f42640299eca308",
			"c94d07416c128e5db78a00323a9b81bedfe3fa5cc97951c906454c24203899c3",
		}},
		{NewTransactionSigningHashWriter(), []string{
			"2abf299e2133a860348dddad77a68cd67dba9f855528184be325fba18ae8faf1",
			"904dbf56bdfcd00ac2d808d0ead76922d361815c1bf16a25ed10de02c10b91e5",
			"859d9d7521422e37ee075e55abd09eba052656848ae395d94fc1613935fb5b58",
			"bb905f5372283590216ec9bad011f2d50951432b9e99ce2d7edc7727011d8da4",
			"3c9fc77e95d1b0cc0c60d1ba34f5efa72bd43f58b2fb95ffcf36c55755393583",
		}},
		{NewTransactionSigningHashECDSAWriter(), []string{
			"b31ad1fbbe41b0e2a90e07c84708b38ba581f0c0e9185416913a04fb6d342027",
			"c43e1f75ea9df6379b56a95074c2b6289ed8c5a01fff2d49d9d44ad5575c164b",
			"49085f99fa0084b5436663f757a5916b1e4290c3321707fb76921ed4e47844ec",
			"3f887e866428de813c1d0463b14eef3ca1363c8187e917dda1eee0ec5996490b",
			"56de89a8c75f0fee2de61b11ab05d0d42e29ed50879467cf128dd80800a52ada",
		}},
		{NewBlockHashWriter(), []string{
			"b5276fb2a7776cbcf5e57e329a840ac2951bc711231ccab04deae865d9519208",
			"19b043f6f44877e3781fe157d50473d2511d5d36349a2b2db678e83489a8c83b",
			"fe41721f230590b47e166c4bff2eceae4903d93f85a02992d558c5f2bf3f3bc9",
			"eae5771e1a28e403af4fd3c3f19d30cf4196afc0f70f89422dec10faf3eb1470",
			"6763451bf7feb77b835af722a27c141e6f41f6cc5b47f2dad02e40094b5c1571",
		}},
		{NewMerkleBranchHashWriter(), []string{
			"ab3321cbc7ad7b002bfae462d80a94c6b66c3b3d22f62b6b0f21362684f3084d",
			"b067715dc4243436afac9d7a087ee3be081b06c3d0429e093bbe6046e79f018e",
			"e7d17e894e90ae8b54c50511291e838bd70d5a35d8060bc5f75b71491a4809d3",
			"457a0288dcd184a0f34a0474c30f5d8aaa5437e37bf6d3b10bd4fe96c0cd705f",
			"67cb6d1507d20bd60f550f0ec228176c57f73581473718b02556997b751e8a83",
		}},
	}

	for _, testVector := range tests {
		hasher := testVector.hasher
		for i, data := range datas {
			hasher.InfallibleWrite(data)
			res := hasher.Finalize().String()
			if res != testVector.expected[i] {
				panic(fmt.Sprintf("expected: %s, got: %s", testVector.expected[i], res))
			}
		}
	}

}

func BenchmarkNewBlockHashWriterSmall(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	var someBytes [32]byte
	r.Read(someBytes[:])
	for i := 0; i < b.N; i++ {
		hasher := NewBlockHashWriter()
		hasher.InfallibleWrite(someBytes[:])
		hasher.Finalize()
	}
}

func BenchmarkNewBlockHashWriterBig(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	var someBytes [1024]byte
	r.Read(someBytes[:])
	for i := 0; i < b.N; i++ {
		hasher := NewBlockHashWriter()
		hasher.InfallibleWrite(someBytes[:])
		hasher.Finalize()
	}
}

func BenchmarkNewHeavyHashWriterSmall(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	var someBytes [32]byte
	r.Read(someBytes[:])
	for i := 0; i < b.N; i++ {
		hasher := NewHeavyHashWriter()
		hasher.InfallibleWrite(someBytes[:])
		hasher.Finalize()
	}
}

func BenchmarkNewHeavyHashWriterBig(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	var someBytes [1024]byte
	r.Read(someBytes[:])
	for i := 0; i < b.N; i++ {
		hasher := NewHeavyHashWriter()
		hasher.InfallibleWrite(someBytes[:])
		hasher.Finalize()
	}
}
