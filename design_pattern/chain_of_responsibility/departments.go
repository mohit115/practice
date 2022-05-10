package main

import "time"

type Departmenti interface {
	setNext(d Departmenti) Departmenti
	execute(p *Patient, update func(id string, val interface{}))
}

type Reception struct {
	next Departmenti
}

func (r *Reception) execute(p *Patient, update func(id string, val interface{})) {
	time.Sleep(time.Millisecond*100)
	p.registrationDone = true
	update(p.id, p)
	r.next.execute(p, update)

}

func (r *Reception) setNext(d Departmenti) Departmenti {
	r.next = d
	return r.next
}

type Doctor struct {
	next Departmenti
}

func (r *Doctor) execute(p *Patient, update func(id string, val interface{})) {
	time.Sleep(time.Millisecond*150)
	p.doctorCheckUpDone = true
	update(p.id, p)
	r.next.execute(p, update)

}

func (r *Doctor) setNext(d Departmenti) Departmenti {
	r.next = d
	return r.next
}

type Medicine struct {
	next Departmenti
}

func (r *Medicine) execute(p *Patient, update func(id string, val interface{})) {
	time.Sleep(time.Millisecond*200)
	p.medicineDone = true
	update(p.id, p)
	r.next.execute(p, update)
}

func (r *Medicine) setNext(d Departmenti) Departmenti {
	r.next = d
	return r.next
}

type Cashier struct {
	next Departmenti
}

func (r *Cashier) execute(p *Patient, update func(id string, val interface{})) {
	time.Sleep(time.Millisecond*250)
	p.paymentDone = true
	update(p.id, p)
}

func (r *Cashier) setNext(d Departmenti) Departmenti {
	r.next = d
	return r.next
}
