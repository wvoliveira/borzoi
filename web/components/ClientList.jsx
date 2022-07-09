import useAuth from "../lib/hooks/useAuth";
import useClients from "../lib/hooks/useClients";


export default function ClientList() {
  const { auth } = useAuth({
    redirectTo: "/login",
  });
  const { clients } = useClients(auth);

  console.log(clients);

  return (
    <div>
    <table>
        <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Action</th>
        </tr>
        {clients?.data?.clients.map((item, index) => {
            return (
                <tr key={index}>
                    <td>{item.name}</td>
                    <td>{item.description}</td>
                    <td>† ¥ ∆ µ π</td>
                </tr>
            )
        })}
    </table>
    </div>
);
}