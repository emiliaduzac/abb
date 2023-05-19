
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

func (abb *Abb[K, V]) buscarPadre(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	act := nodo
	if act.izquierdo == nil && act.derecho == nil { // si llega a este caso es xq no encontre el nodo al cual le busco padre
		return nil
	} else if act.izquierdo.clave == clave || act.derecho.clave == clave { // si el hijo de act == clave, encontre al padre
		return act
	}

	if abb.cmp(clave, act.clave) > 0 {
		act = abb.buscarPadre(act.derecho, clave) //si la clave que busco es mayor a la actual, busco del lado derecho
	} else {
		act = abb.buscarPadre(act.izquierdo, clave) //si la clave que busco es menor a la actual, busco del lado izq
	}

	return act
}

func (abb *Abb[K, V]) recorrerInOrder(nodo nodoAbb[K, V]) nodoAbb[K, V] {
	var ant *nodoAbb[K, V]
	act := nodo.derecho
	for act != nil {
		ant, act = act, act.izquierdo
	}

	return *ant
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

func (abb *Abb[K, V]) Borrar(clave K) V {
	nodoBuscado := abb.buscarHoja(nil, abb.raiz, clave)
	nodoPadre := abb.buscarPadre(abb.raiz, clave)

	if nodoBuscado.derecho == nil && nodoBuscado.izquierdo == nil { // CASO SIN HIJOS
		if abb.cmp(nodoBuscado.clave, nodoPadre.clave) > 0 {
			nodoPadre.derecho = nil
		} else {
			nodoPadre.izquierdo = nil
		}

	} else if nodoBuscado.derecho == nil && nodoBuscado.izquierdo != nil { // CASO UN HIJO IZQ
		if abb.cmp(nodoBuscado.clave, nodoPadre.clave) > 0 {
			nodoPadre.derecho = nodoBuscado.izquierdo
		} else {
			nodoPadre.izquierdo = nodoBuscado.izquierdo
		}
	} else if nodoBuscado.derecho != nil && nodoBuscado.izquierdo == nil { // CASO UN HIJO DER
		if abb.cmp(nodoBuscado.clave, nodoPadre.clave) > 0 {
			nodoPadre.derecho = nodoBuscado.derecho
		} else {
			nodoPadre.izquierdo = nodoBuscado.derecho
		}

	} else { // CASO DOS HIJOS
		reemplazante := abb.recorrerInOrder(*nodoBuscado)
		if abb.cmp(nodoBuscado.clave, nodoPadre.clave) > 0 {
			nodoPadre.derecho = &reemplazante
		} else {
			nodoPadre.izquierdo = &reemplazante
		}
		abb.Borrar(reemplazante.clave)
	}
	return nodoBuscado.dato
}

func (abb Abb[K, V]) Obtener(clave K) V {
	valor := abb.buscarHoja(nil, abb.raiz, clave)

	if !abb.Pertenece(valor.clave) {
		panic("La clave no pertenece al diccionario")
	}

	return valor.dato
}

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
	cmp       func(K, K) int
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

// iteradores por rango
func (abb Abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}

	if abb.cmp(abb.raiz.clave, *desde) > 0 {
		abb.raiz.izquierdo.iterarNodo(visitar)
	}

	if abb.cmp(abb.raiz.clave, *desde) > 0 && abb.cmp(abb.raiz.clave, *hasta) < 0 {
		abb.raiz.iterarNodo(visitar)
	}

	if abb.cmp(abb.raiz.clave, *hasta) < 0 {
		abb.raiz.derecho.iterarNodo(visitar)
	}

}

func (abb Abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterRango := new(IteradorDiccionarioABB[K, V])
	pila := pila.CrearPilaDinamica[nodoAbb[K, V]]()
	funcionCmp := abb.cmp

	iterRango.pilaNodos = pila
	iterRango.cmp = funcionCmp

	if abb.cmp(abb.raiz.clave, *desde) > 0 {
		iterRango.apilarHijosIzquierdos(abb.raiz)
	}

	return iterRango
}

func (iter IteradorDiccionarioABB[K, V]) VerActualIterRango(desde *K, hasta *K) (K, V) {
	return iter.pilaNodos.VerTope().clave, iter.pilaNodos.VerTope().dato
}

func (iter IteradorDiccionarioABB[K, V]) SiguienteIterRango(desde *K, hasta *K) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	act := iter.pilaNodos.Desapilar()
	if iter.cmp(act.clave, *desde) > 0 && act.derecho != nil {
		iter.apilarHijosIzquierdos(act.derecho)
	}
}

func (iter IteradorDiccionarioABB[K, V]) HaySiguienteIterRango(desde *K, hasta *K) bool {
	if iter.cmp(iter.pilaNodos.VerTope().clave, *hasta) > 0 {
		return false
	}

	return iter.pilaNodos.EstaVacia()
}
