package diccionario

type DiccionarioOrdenado[K comparable, V any] interface {
	// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
	Guardar(clave K, dato V)

	// Pertenece determina si una clave ya se encuentra en el diccionario, o no
	Pertenece(clave K) bool

	// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
	// 'La clave no pertenece al diccionario'
	Obtener(clave K) V

	// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
	// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
	Borrar(clave K) V

	// Cantidad devuelve la cantidad de elementos dentro del diccionario
	Cantidad() int

	// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
	// mismo
	Iterar(func(clave K, dato V) bool)

	//Iterador devuelve un IterDiccionario para este Diccionario
	Iterador() IterDiccionario[K, V]

	// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
	// incluyéndolos en caso de encontrarse
	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)

	// IteradorRango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
