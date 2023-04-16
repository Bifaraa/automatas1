import './index.css'
import { Consola } from './components/Consola'
import { Graphviz } from 'graphviz-react'
import { useEffect, useState } from 'react'
import { useDotfetch } from './hooks/useDotfetch'

import { Transition } from 'react-transition-group'
import { useRef } from 'react'

const duration = 300

const defaultStyle = {
  transition: `opacity ${duration}s ease-in-out 1s`,
  opacity: 0,
}

const transitionStyles = {
  entering: { opacity: 1 },
  entered: { opacity: 1 },
  exiting: { opacity: 0 },
  exited: { opacity: 0 },
}

function App() {
  const [loading, setLoading] = useState(true)
  const { dot, dotFetchUpdate } = useDotfetch()

  const [inProp, setInProp] = useState(false)
  const nodeRef = useRef(null)

  const handdleClic = () => {
    dotFetchUpdate
    setLoading(false)
    setInProp(true)
  }
  return (
    <>
      <header className='w-screen h-[15vh] bg-[#1B2430] shadow-lg shadow-black/30 flex justify-center'>
        <h1 className='font-mono text-[2.5em] text-center p-5 text-white '>
          Automata (AFND)
        </h1>
      </header>
      <main className='w-screen min-h-[85vh] flex justify-center relative'>
        {loading ? (
          <div
            className='relative text-[3em] cursor-pointer p-10 rounded-3xl shadow-[10px_10px_10px_rgba(0,0,0,0.5)] self-center bg-[#D6D5A8] '
            onClick={handdleClic}
          >
            <span className='after:content-[""] after:absolute after:right-0 after:bottom-0 after:block after:w-0 after:h-full after:bg-[#afae81]/30 hover:after:w-full z-10 after:rounded-3xl after:transition-all after:duration-1000'>
              <a className='relative z-40 ' href='#'>
                Iniciar Automata
              </a>
            </span>
          </div>
        ) : (
          <Transition
            nodeRef={nodeRef}
            in={inProp}
            timeout={500}
            unmountOnExit
            onEnter={() => setShowButton(false)}
            onExited={() => setShowButton(true)}
          >
            {(state) => (
              <div
                className='w-full flex justify-between '
                ref={nodeRef}
                style={{
                  ...defaultStyle,
                  ...transitionStyles[state],
                }}
              >
                <section className='ml-12'>
                  <Graphviz dot={`${dot}`} />
                </section>

                <Consola />
              </div>
            )}
          </Transition>
        )}
      </main>
    </>
  )
}

export default App
