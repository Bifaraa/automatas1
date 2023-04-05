import { useEffect, useState } from 'react'

export function TypingEffect({ text, speed }) {
  const [displayText, setDisplatText] = useState('')
  const [currentIndex, setCurrentIndex] = useState(0)
  const [showCursor, setShowCursor] = useState(true)

  useEffect(() => {
    if (!text.length) return

    const randomTime = Math.floor(Math.random() * speed) + 200

    const intervalID = setInterval(() => {
      if (currentIndex >= text.length) {
        clearInterval(intervalID)
        setShowCursor(false)
        return
      }
      const nextIndex = text.indexOf(' ', currentIndex + 1)
      if (nextIndex < 0) {
        setDisplatText(text)
        setCurrentIndex(text.length)
        return
      }

      setDisplatText(text.slice(0, nextIndex))
      setCurrentIndex(currentIndex + 1)
    }, randomTime)

    return () => clearInterval(intervalID)
  }, [text, currentIndex])

  return (
    <span
      className={`${
        showCursor ? 'after:content-["▋"] after:ml-1 after:animate-typing' : ''
      }`}
    >{`----> ${displayText}`}</span>
  )
}
