package env_test

import (
	"errors"
	"github.com/darkhelmet/env"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestStrings(c *C) {
	c.Assert(func() { env.String("DOESNT_EXIST") }, Panics, errors.New("env: Environment variable DOESNT_EXIST doesn't exist"))

	os.Setenv("test", "gocheck")
	test := map[string]string{
		env.String("test"):                                               "gocheck",
		env.StringDefault("foobar", "fizzbin"):                           "fizzbin",
		env.StringDefaultF("foobar", func() string { return "fizzbot" }): "fizzbot",
	}

	for obtained, expected := range test {
		c.Assert(obtained, Equals, expected)
	}
}

func (s *S) TestInts(c *C) {
	c.Assert(func() { env.Int("DOESNT_EXIST") }, Panics, errors.New("env: Environment variable DOESNT_EXIST doesn't exist"))

	os.Setenv("test", "a")
	c.Assert(func() { env.Int("test") }, Panics, errors.New(`env: failed parsing int: strconv.ParseInt: parsing "a": invalid syntax`))

	os.Setenv("test", "1")
	os.Setenv("test2", "02")
	test := map[int]int{
		env.Int("test"):                                    1,
		env.Int("test2"):                                   2,
		env.IntDefault("foobar", 3):                        3,
		env.IntDefaultF("foobar", func() int { return 4 }): 4,
	}

	for obtained, expected := range test {
		c.Assert(obtained, Equals, expected)
	}
}

func (s *S) TestFloats(c *C) {
	c.Assert(func() { env.Float("DOESNT_EXIST") }, Panics, errors.New("env: Environment variable DOESNT_EXIST doesn't exist"))

	os.Setenv("test", "a")
	c.Assert(func() { env.Float("test") }, Panics, errors.New(`env: failed parsing float: strconv.ParseFloat: parsing "a": invalid syntax`))

	os.Setenv("test", "1.1")
	os.Setenv("test2", "02.2")
	test := map[float64]float64{
		env.Float("test"):                                          1.1,
		env.Float("test2"):                                         2.2,
		env.FloatDefault("foobar", 3.3):                            3.3,
		env.FloatDefaultF("foobar", func() float64 { return 4.4 }): 4.4,
	}

	for obtained, expected := range test {
		c.Assert(obtained, Equals, expected)
	}
}

func (s *S) TestBool(c *C) {
	c.Assert(func() { env.Bool("DOESNT_EXIST") }, Panics, errors.New("env: Environment variable DOESNT_EXIST doesn't exist"))

	os.Setenv("test", "a")
	c.Assert(func() { env.Bool("test") }, Panics, errors.New(`env: failed parsing boolean: strconv.ParseBool: parsing "a": invalid syntax`))

	os.Setenv("test", "1")
	os.Setenv("test2", "0")
	os.Setenv("test3", "t")
	os.Setenv("test4", "f")
	os.Setenv("test5", "T")
	os.Setenv("test6", "F")
	os.Setenv("test7", "TRUE")
	os.Setenv("test8", "FALSE")
	os.Setenv("test9", "true")
	os.Setenv("test10", "false")
	os.Setenv("test11", "True")
	os.Setenv("test12", "False")

	test := map[bool]bool{
		env.Bool("test"):                true,
		env.Bool("test2"):               false,
		env.Bool("test3"):               true,
		env.Bool("test4"):               false,
		env.Bool("test5"):               true,
		env.Bool("test6"):               false,
		env.Bool("test7"):               true,
		env.Bool("test8"):               false,
		env.Bool("test9"):               true,
		env.Bool("test10"):              false,
		env.Bool("test11"):              true,
		env.Bool("test12"):              false,
		env.BoolDefault("foobar", true): true,
	}

	for obtained, expected := range test {
		c.Assert(obtained, Equals, expected)
	}
}
