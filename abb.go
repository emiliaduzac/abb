package diccionario

import "tdas/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type Abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dic := new(Abb[K, V])
	return dic
}

func crearHoja[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nuevaHoja := new(nodoAbb[K, V])
	nuevaHoja.clave = clave
	nuevaHoja.dato = dato
	return nuevaHoja
}

func (abb *Abb[K, V]) buscarHoja(nodoPadre *nodoAbb[K, V], nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	act := nodo
	if act == nil {
		return nodoPadre
	} else if act.clave == clave {
		return act
	}

	if abb.cmp(clave, act.clave) > 0 {
		act = abb.buscarHoja(act, act.derecho, clave)
	} else {
		act = abb.buscarHoja(act, act.izquierdo, clave)
	}

	return act
}

func (abb *Abb[K, V]) Guardar(clave K, dato V) {
	hoja := crearHoja(clave, dato)

	if abb.cantidad == 0 {
		abb.raiz = hoja
	}

	posicion := abb.buscarHoja(nil, abb.raiz, clave)

	if posicion.clave == clave {
		posicion = hoja
	}

	if abb.cmp(clave, posicion.clave) > 1 {
		posicion.derecho = hoja
	} else {
		posicion.izquierdo = hoja
	}

	abb.cantidad++
}

func (abb Abb[K, V]) Pertenece(clave K) bool {
	posicion := abb.buscarHoja(nil, abb.raiz, clave)

	return posicion.clave == clave
}

// func (abb *Abb[K, V]) Borrar(clave K, dato V) V                                          {}

// func (abb Abb[K, V]) Obtener(clave K) V                                                  {}

func (abb Abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *Abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iterarNodo(visitar)
}

func (nodo *nodoAbb[K, V]) iterarNodo(visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	nodo.izquierdo.iterarNodo(visitar)
	visitar(nodo.clave, nodo.dato)
	nodo.derecho.iterarNodo(visitar)
}

type IteradorDiccionarioABB[K comparable, V any] struct {
	pilaNodos pila.Pila[nodoAbb[K, V]]
}

func (abb Abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(IteradorDiccionarioABB[K, V])
	pila := pila.CrearPilaDinamica[nodoAbb[K, V]]()

	iter.pilaNodos = pila

	iter.apilarHijosIzquierdos(abb.raiz)

	return iter
}

func (iter *IteradorDiccionarioABB[K, V]) apilarHijosIzquierdos(nodo *nodoAbb[K, V]) {
	act := nodo
	iter.pilaNodos.Apilar(*act)
	for act != nil {
		iter.pilaNodos.Apilar(*act.izquierdo)
		act = act.izquierdo
	}
}

func (iter IteradorDiccionarioABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.pilaNodos.VerTope().clave, iter.pilaNodos.VerTope().dato
}

func (iter *IteradorDiccionarioABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	act := iter.pilaNodos.Desapilar()
	iter.apilarHijosIzquierdos(act.derecho)
}

func (iter IteradorDiccionarioABB[K, V]) HaySiguiente() bool {
	return iter.pilaNodos.EstaVacia()
}

// func (abb Abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {}
// func (abb Abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]             {}
