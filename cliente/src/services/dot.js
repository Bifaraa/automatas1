export const dotFetch = async () => {
  const res = await fetch('http://localhost:3000/dot')
  const dat = await res.json()
  return dat.data
}
