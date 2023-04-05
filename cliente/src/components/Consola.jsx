import { useState, useEffect } from 'react'
import data from '../../../entradas.json'
import { TypingEffect } from './TypingEffect'

export function Consola() {
  //TODO QUE LEA EL MISMO JSON QUE EL BACKEND
  const [respuesta, setRespuesta] = useState([])
  const [visible, setVisible] = useState(false)
  const [num, setNum] = useState('')

  useEffect(() => {
    const fetchRespuesta = async () => {
      const response = await fetch('http://localhost:3000/run')
      const { data } = await response.json()
      setRespuesta(data)
    }
    fetchRespuesta()
  }, [])

  const renderData = respuesta.map((item, index) => (
    <li key={index}>
      <TypingEffect text={item} speed={80} />
    </li>
  ))

  const handdleClic = () => {
    setVisible(!visible)
  }

  const atras = () => {
    if (visible === false) return console.log('no me puedo devolver')
    console.log('me devolvi')
    setVisible(true)
  }

  const handleSubmit = async (event) => {
    event.preventDefault()
    console.log(num)
    const res = await fetch('http://localhost:3000/ejecutar', {
      method: 'POST',
      body: JSON.stringify({ num: num }),
    })
    const data = await res.json()
    console.log(data)
  }

  return (
    <section className='w-[50%] min-h-[100%] flex justify-center mt-10 font-mono text-white '>
      <article className='bg-[#816797]  w-[80%] max-h-[60%] min-h-[40%] flex items-start flex-col gap-4 rounded-lg shadow-xl shadow-black/40 overflow-auto '>
        <div className='w-full flex justify-center p-8 bg-[#51557E] rounded-lg shadow-sm relative'>
          <h2 className='text-2xl text-center '>Consola</h2>
          <div onClick={atras} className='absolute left-10 cursor-pointer'>
            <svg
              xmlns='http://www.w3.org/2000/svg'
              width='30'
              height='30'
              fill='currentColor'
              classNanme='bi bi-arrow-90deg-left'
              viewBox='0 0 16 16'
            >
              <path
                fill-rule='evenodd'
                d='M1.146 4.854a.5.5 0 0 1 0-.708l4-4a.5.5 0 1 1 .708.708L2.707 4H12.5A2.5 2.5 0 0 1 15 6.5v8a.5.5 0 0 1-1 0v-8A1.5 1.5 0 0 0 12.5 5H2.707l3.147 3.146a.5.5 0 1 1-.708.708l-4-4z'
              />
            </svg>
          </div>
        </div>

        {visible ? (
          <>
            <p className='text-xl text-center w-full'>
              Valores Evaluados por defecto:
            </p>
            <ul className='flex flex-col px-6 gap-6'>
              {data.list.map((item, index) => (
                <li key={index}>
                  <TypingEffect text={item.valor} speed={700} />{' '}
                </li>
              ))}
            </ul>
            <h2 className='text-2xl w-full mt-5 text-center'>Salidas</h2>
            <hr className='w-[90%] ' />
            <p className='text-xl text-center w-full'>
              Salida de los valores evaluados
            </p>
            <ul className='flex flex-col gap-6 px-6'>{renderData}</ul>
          </>
        ) : (
          <>
            <form
              onSubmit={handleSubmit}
              className='w-[100%] flex flex-col items-center gap-5 text-black '
            >
              <input
                onChange={(e) => setNum(e.target.value)}
                className='h-10 text-center'
                type='text'
              />
              <button className='text-md bg-[#D6D5A8] rounded-md p-2 shadow-md shadow-black/30'>
                Ingresar Un valor
              </button>
            </form>
            <button
              className='self-center text-md bg-[#D6D5A8] rounded-md text-black p-2 shadow-md shadow-black/30'
              onClick={handdleClic}
            >
              Cargar ejemplo
            </button>
          </>
        )}
      </article>
    </section>
  )
}
