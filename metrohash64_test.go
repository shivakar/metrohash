package metrohash_test

import (
	"testing"

	"github.com/shivakar/metrohash"
	"github.com/stretchr/testify/assert"
)

// Test data
var data = [...]struct {
	input   string
	intHash uint64
	strHash string
}{
	{"", 8097384203561113213, "705fb008071e967d"},
	{"a", 12641362448282233803, "af6f242b7ed32bcb"},
	{"ab", 14784933243920743832, "cd2ea2738fc27d98"},
	{"abc", 17099979927131455419, "ed4f5524e6faffbb"},
	{"abcd", 13572854355886132194, "bc5c771fe6b4cbe2"},
	{"abcde", 15625493671854626812, "d8d8e7b5070257fc"},
	{"abcdef", 3528379674302886064, "30f7529e9f8830b0"},
	{"abcdefg", 9195229972332259068, "7f9c04f100e3a6fc"},
	{"abcdefgh", 8690139311311089811, "789993a94b5af093"},
	{"abcdefghi", 12275996829284600609, "aa5d1a1341548b21"},
	{"abcdefghij", 2464403427817883015, "223351dc52570587"},
	{"abcdefghijk", 15273828071406370431, "d3f789b4ec90fe7f"},
	{"abcdefghijkl", 3784547043968092435, "3485697c2f790913"},
	{"abcdefghijklm", 6394046173423095729, "58bc371e1c544fb1"},
	{"abcdefghijklmn", 13331640262538197665, "b903802f49e27aa1"},
	{"abcdefghijklmno", 856108144412802574, "0be181d24d0b3e0e"},
	{"abcdefghijklmnop", 8034916556097801260, "6f81c209761d4c2c"},
	{"abcdefghijklmnopq", 8293923542328109661, "7319ef84c6b76a5d"},
	{"abcdefghijklmnopqr", 9458067369213740671, "8341ce294bb2c67f"},
	{"abcdefghijklmnopqrs", 13424278878677786681, "ba4c9e83e95e6039"},
	{"abcdefghijklmnopqrst", 9878707767057859577, "8918385fdf6043f9"},
	{"abcdefghijklmnopqrstu", 6906997572826963590, "5fda95b298d61e86"},
	{"abcdefghijklmnopqrstuv", 5950688388136563895, "5295178f97ebf4b7"},
	{"abcdefghijklmnopqrstuvw", 4438600239575069299, "3d991366b3d5f273"},
	{"abcdefghijklmnopqrstuvwx", 5073188467988281691, "46679608610cb95b"},
	{"abcdefghijklmnopqrstuvwxy", 7312434858287640440, "657afcc2bb6b2f78"},
	{"abcdefghijklmnopqrstuvwxyz", 1767508563557181619, "188773ac844968b3"},
	{"1", 400046725477776297, "058d406095a93ba9"},
	{"12", 11768182440106335391, "a350fb942a871c9f"},
	{"123", 8869102568122554874, "7b1561cb917319fa"},
	{"12345", 1542658667971767033, "15689fe29a226af9"},
	{"123456", 1171009367469622072, "104042d0c0ae8338"},
	{"1234567", 6915207505515058674, "5ff7c096825715f2"},
	{"123456789", 12126722693468943237, "a84ac60a21e27385"},
	{"Hello, World!!", 2952858128079109675, "28faa8d29447be2b"},
	{"How are you doing?", 4018291871582581493, "37c3d72ad984f2f5"},
	{"Discard medicine more than two years old.", 4806600389108629132, "42b479966212fe8c"},
	{"He who has a shady past knows that nice guys finish last.", 18116113071188485382, "fb695cdadf002506"},
	{"I wouldn't marry him with a ten foot pole.", 3807230084335824214, "34d5ff9713046d56"},
	{"Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", 9286414470680077242, "80dff8c271eee3ba"},
	{"The days of the digital watch are numbered.  -Tom Stoppard", 13772500942842091206, "bf21c0a33ed5d6c6"},
	{"Nepal premier won't resign.", 11924647779184696464, "a57cdbf9c5f45c90"},
	{"For every action there is an equal and opposite government program.", 8100009507500453014, "706903bb3bd9ec96"},
	{"His money is twice tainted: 'taint yours and 'taint mine.", 13805151511692099319, "bf95c02834dabef7"},
	{"There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", 2949067878655966973, "28ed319c5dcdbefd"},
	{"It's a tiny change to the code and not completely disgusting. - Bob Manchek", 18341731511480495064, "fe8aeba18b622bd8"},
	{"size:  a.out:  bad magic", 986235606592326546, "0dafd00f13c70b92"},
	{"The major problem is with sendmail.  -Mark Horton", 9679746698404454194, "86555e5632285f32"},
	{"Give me a rock, paper and scissors and I will move the world.  CCFestoon", 1447506096505631214, "14169320246e31ee"},
	{"If the enemy is within range, then so are you.", 8344747233677348917, "73ce7f658bc31035"},
	{"It's well we cannot hear the screams/That we create in others' dreams.", 2117962786679881411, "1d6483eedb8cfec3"},
	{"You remind me of a TV show, but that's all right: I watch it anyway.", 4651850682360023146, "408eb18c8e3a486a"},
	{"C is as portable as Stonehedge!!", 4690659838578788536, "4118924560ae08b8"},
	{"Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", 4808728188833102663, "42bc08cf5ccdfb47"},
	{"The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", 9789315181472652520, "87daa24aa2041ce8"},
	{"How can you write a big system without C++?  -Paul Glick", 17926697427089629216, "f8c86c5489962020"},
}

func Test_NewMetroHash64(t *testing.T) {
	assert := assert.New(t)

	r1 := new(metrohash.MetroHash64)
	r1.Reset()
	r2 := metrohash.NewMetroHash64()
	r3 := metrohash.NewSeedMetroHash64(0)

	assert.Equal(r1, r2)
	assert.Equal(r2, r3)

	// Size and BlockSize
	assert.Equal(r1.Size(), 8)
	assert.Equal(r3.Size(), r1.Size())
	assert.Equal(r1.BlockSize(), 32)
	assert.Equal(r2.BlockSize(), r1.BlockSize())
}

func Test_Hash(t *testing.T) {
	assert := assert.New(t)
	r := metrohash.NewMetroHash64()
	for _, s := range data {
		r.Reset()
		r.Write([]byte(s.input))
		assert.Equal(s.intHash, r.Uint64())
		assert.Equal(s.strHash, r.String())
	}
}

func Test_StreamingWrite(t *testing.T) {
	assert := assert.New(t)

	input := "abcdefghijklmnopqrstuvwxyz0123456789"

	r1 := metrohash.NewMetroHash64()
	r2 := metrohash.NewMetroHash64()

	for i, v := range []byte(input) {
		r1.Reset()
		r1.Write([]byte(input[:i+1]))
		r2.Write([]byte(string(v)))

		assert.Equal(r1.Uint64(), r2.Uint64())
		assert.Equal(r1.String(), r2.String())
	}
}
