export const Fetcher = (...args) => fetch(...args).then((res) => res.json())
