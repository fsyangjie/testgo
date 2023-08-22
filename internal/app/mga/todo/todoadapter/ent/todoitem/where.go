// Code generated by entc, DO NOT EDIT.

package todoitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Completed applies equality check predicate on the "completed" field. It's identical to CompletedEQ.
func Completed(v bool) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUID), v))
	})
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUID), v))
	})
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUID), v))
	})
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUID), v))
	})
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// CompletedEQ applies the EQ predicate on the "completed" field.
func CompletedEQ(v bool) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// CompletedNEQ applies the NEQ predicate on the "completed" field.
func CompletedNEQ(v bool) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompleted), v))
	})
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrder), v))
	})
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrder), v...))
	})
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrder), v...))
	})
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrder), v))
	})
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrder), v))
	})
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrder), v))
	})
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrder), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TodoItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TodoItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TodoItem) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TodoItem) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TodoItem) predicate.TodoItem {
	return predicate.TodoItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
