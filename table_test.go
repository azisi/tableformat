package tableformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableSort(t *testing.T) {
	testData := []struct {
		header   []string
		orderby  []int
		unsorted [][]string
		sorted   [][]string
	}{
		{
			header:  []string{"Surname", "Name", "Team", "Country"},
			orderby: []int{0, 1},
			unsorted: [][]string{
				{"ANTETOKOUNMPO", "Thanasis", "Milwaukee Bucks", "USA", "1992"},
				{"ANTETOKOUNMPO", "Giannis", "Milwaukee Bucks", "USA", "1994"},
				{"PAPANIKOLAOU", "Konstantinos", "Olympiacos", "GRE", "1990"},
				{"PAPAPETROU", "Ioannis", "Panathinaikos", "GRE", "1994"},
				{"PRINTEZIS", "Georgios", "Olympiacos", "GRE", "1985"},
				{"SLOUKAS", "Konstantinos", "FENERBAHCE", "TUR", "1990"},
				{"VASILOPOULOS", "Panagiotis", "Peristeri", "GRE", "1984"},
				{"PAPAGIANNIS", "Georgios", "Panathinaikos", "GRE", "1997"},
				{"MANTZARIS", "Evangelos", "UNICS Kazan", "RUS", "1990"},
				{"LARENTZAKIS", "Giannoulis", "AEK", "GRE", "1993"},
				{"KONIARIS", "Antonis", "Olympiacos", "GRE", "1997"},
				{"BOUROUSIS", "Ioannis", "Zhejiang Lions", "CHN", "1983"},
				{"CALATHES", "Nick", "Panathinaikos", "GRE", "1989"},
			},
			sorted: [][]string{
				{"ANTETOKOUNMPO", "Giannis", "Milwaukee Bucks", "USA", "1994"},
				{"ANTETOKOUNMPO", "Thanasis", "Milwaukee Bucks", "USA", "1992"},
				{"BOUROUSIS", "Ioannis", "Zhejiang Lions", "CHN", "1983"},
				{"CALATHES", "Nick", "Panathinaikos", "GRE", "1989"},
				{"KONIARIS", "Antonis", "Olympiacos", "GRE", "1997"},
				{"LARENTZAKIS", "Giannoulis", "AEK", "GRE", "1993"},
				{"MANTZARIS", "Evangelos", "UNICS Kazan", "RUS", "1990"},
				{"PAPAGIANNIS", "Georgios", "Panathinaikos", "GRE", "1997"},
				{"PAPANIKOLAOU", "Konstantinos", "Olympiacos", "GRE", "1990"},
				{"PAPAPETROU", "Ioannis", "Panathinaikos", "GRE", "1994"},
				{"PRINTEZIS", "Georgios", "Olympiacos", "GRE", "1985"},
				{"SLOUKAS", "Konstantinos", "FENERBAHCE", "TUR", "1990"},
				{"VASILOPOULOS", "Panagiotis", "Peristeri", "GRE", "1984"},
			},
		},
		{
			header:  []string{"Surname", "Name", "Team", "Country", "Born"},
			orderby: []int{3, 2, 4},
			unsorted: [][]string{
				{"ANTETOKOUNMPO", "Giannis", "Milwaukee Bucks", "USA", "1994"},
				{"ANTETOKOUNMPO", "Thanasis", "Milwaukee Bucks", "USA", "1992"},
				{"BOUROUSIS", "Ioannis", "Zhejiang Lions", "CHN", "1983"},
				{"CALATHES", "Nick", "Panathinaikos", "GRE", "1989"},
				{"PAPANIKOLAOU", "Konstantinos", "Olympiacos", "GRE", "1990"},
				{"PAPAPETROU", "Ioannis", "Panathinaikos", "GRE", "1994"},
				{"PRINTEZIS", "Georgios", "Olympiacos", "GRE", "1985"},
				{"SLOUKAS", "Konstantinos", "Fenerbahce", "TUR", "1990"},
				{"VASILOPOULOS", "Panagiotis", "Peristeri", "GRE", "1984"},
				{"PAPAGIANNIS", "Georgios", "Panathinaikos", "GRE", "1997"},
				{"MANTZARIS", "Evangelos", "Unics Kazan", "RUS", "1990"},
				{"LARENTZAKIS", "Giannoulis", "AEK", "GRE", "1993"},
				{"KONIARIS", "Antonis", "Olympiacos", "GRE", "1997"},
			},
			sorted: [][]string{
				{"BOUROUSIS", "Ioannis", "Zhejiang Lions", "CHN", "1983"},
				{"LARENTZAKIS", "Giannoulis", "AEK", "GRE", "1993"},
				{"PRINTEZIS", "Georgios", "Olympiacos", "GRE", "1985"},
				{"PAPANIKOLAOU", "Konstantinos", "Olympiacos", "GRE", "1990"},
				{"KONIARIS", "Antonis", "Olympiacos", "GRE", "1997"},
				{"CALATHES", "Nick", "Panathinaikos", "GRE", "1989"},
				{"PAPAPETROU", "Ioannis", "Panathinaikos", "GRE", "1994"},
				{"PAPAGIANNIS", "Georgios", "Panathinaikos", "GRE", "1997"},
				{"VASILOPOULOS", "Panagiotis", "Peristeri", "GRE", "1984"},
				{"MANTZARIS", "Evangelos", "Unics Kazan", "RUS", "1990"},
				{"SLOUKAS", "Konstantinos", "Fenerbahce", "TUR", "1990"},
				{"ANTETOKOUNMPO", "Thanasis", "Milwaukee Bucks", "USA", "1992"},
				{"ANTETOKOUNMPO", "Giannis", "Milwaukee Bucks", "USA", "1994"},
			},
		},
	}

	for _, data := range testData {

		tst := NewTable(data.header)
		tst.AppendBulk(data.unsorted)
		tst.OrderBy(data.orderby)
		tst.Sort()
		assert := assert.New(t)
		assert.EqualValues(tst.rows, data.sorted)
		tst.Print()
	}
}
