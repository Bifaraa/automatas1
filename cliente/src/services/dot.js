export const dotFetch = async () => {
  const res = await fetch('/dot')
  const dat = await res.json()
  return dat.data
}
