import useSWR from "swr";
import fetchJson from "./fetchJson";


export default function useClients(auth) {
  const { data: clients } = useSWR(auth ? `/api/v1/clients` : null, fetchJson);
  return { clients };
}
