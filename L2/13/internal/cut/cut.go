package cut

import (
	"bufio"
	"cut/internal/config"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Cut struct {
	cfg *config.Config
	srs io.Reader
}

func New(config *config.Config, srs io.Reader) *Cut {
	return &Cut{
		cfg: config,
		srs: srs,
	}
}

func (c *Cut) DoCut() error {
	fields, err := ParseFields(c.cfg.Fields)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(c.srs)

	for scanner.Scan() {
		line := scanner.Text()

		if c.cfg.Separated && !strings.Contains(line, c.cfg.Delimiter) {
			continue
		}

		sepLine := strings.Split(line, c.cfg.Delimiter)

		var out []string

		for i := 0; fields[i] < len(sepLine); i++ {
			out = append(out, sepLine[fields[i]-1])
		}
		fmt.Println(strings.Join(out, c.cfg.Delimiter))
	}
	return nil
}

func ParseFields(fields string) ([]int, error) {
	if strings.Contains(fields, "0") {
		return []int{}, fmt.Errorf("values may not include zero")
	}

	resFields := make([]int, 0)

	parts := strings.Split(fields, ",")

	for _, p := range parts {
		if strings.Contains(p, "-") {
			ranges := strings.Split(p, "-")
			if len(ranges) == 2 {
				start, err := strconv.Atoi(ranges[0])
				if err != nil {
					return []int{}, fmt.Errorf("illegal list value %v", err)
				}
				end, err := strconv.Atoi(ranges[1])
				if err != nil {
					return []int{}, fmt.Errorf("illegal list value %v", err)
				}

				for i := start; i <= end; i++ {
					resFields = append(resFields, i)
				}
			}
		} else {
			num, err := strconv.Atoi(p)
			if err != nil {
				return []int{}, fmt.Errorf("illegal list value %v", err)
			}
			resFields = append(resFields, num)
		}
	}

	return resFields, nil
}
