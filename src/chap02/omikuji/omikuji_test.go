/* テストの実行方法
特定のパッケージのテスト

- コマンド

```sh
cd src/chap02
go test -v ./omikuji
```

- 出力

```sh
=== RUN   TestOmikuji
--- PASS: TestOmikuji (0.00s)
PASS
ok      chap02/omikuji  0.412s
```

*/

package omikuji

import (
	"testing"
)

func TestOmikuji(t *testing.T) {
	patterns := []struct {
		dice_number int
		expected string
	}{
		{1, "凶"},
		{2, "吉"},
		{3, "吉"},
		{4, "中吉"},
		{5, "中吉"},
		{6, "大吉"},
	}

	for idx, pattern := range patterns {
		actual := Omikuji(pattern.dice_number)
		if pattern.expected != actual {
			t.Errorf(("pattern %d: want %s, actual %s"), idx, pattern.expected, actual)
		}
	}
}