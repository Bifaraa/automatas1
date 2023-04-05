import { useState, useEffect } from 'react'
import { dotFetch } from '../services/dot'

export function useDotfetch() {
  const [dot, setDot] = useState()
  console.log('Using dot')
  const dotFetchUpdate = () => {
    dotFetch().then((newDot) => {
      setDot(newDot)
      console.log(newDot)
    })
  }
  useEffect(dotFetchUpdate, [])
  return { dot, dotFetchUpdate }
}
