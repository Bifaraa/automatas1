import { Graphviz } from 'graphviz-react'
import { useEffect, useState } from 'react'
import { useDotfetch } from '../hooks/useDotfetch'

export function Grafico() {
  const [loading, setLoading] = useState(true)
  const { dot, dotFetchUpdate } = useDotfetch()

  const handdleClic = () => {
    dotFetchUpdate
    setLoading(false)
  }

  return (
    <>
      {loading ? (
        <div onClick={handdleClic}>
          <span>Cargar Grafo</span>
        </div>
      ) : (
        <section className=''>
          <Graphviz dot={`${dot}`} />:
        </section>
      )}
    </>
  )
}
