package oauth2

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestACMemoryStore(t *testing.T) {
	Convey("AC memory store test", t, func() {
		store := NewACMemoryStore(1)
		item := ACInfo{
			ClientID:  "123456",
			UserID:    "999999",
			CreateAt:  time.Now().Unix(),
			ExpiresIn: time.Millisecond * 500,
		}

		Convey("Put Test", func() {
			id, err := store.Put(item)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			Convey("Take Test", func() {
				info, err := store.TakeByID(id)
				So(err, ShouldBeNil)
				So(info.ClientID, ShouldEqual, item.ClientID)
				So(info.UserID, ShouldEqual, item.UserID)
			})
		})

		Convey("GC Test", func() {
			id, err := store.Put(item)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			Convey("Take GC Test", func() {
				time.Sleep(time.Millisecond * 1500)
				info, err := store.TakeByID(id)
				So(err, ShouldNotBeNil)
				So(info, ShouldBeNil)
			})
		})
	})
}
