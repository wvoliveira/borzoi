import useUser from "../lib/hooks/useUser";


export default function ProfileInfo() {
    const { user } = useUser({
        redirectTo: "/login",
    });

    console.log(user);

    if (!user) {
        return (
            <p>Loading...</p>
        )
    }

    console.log(user);

    return (
        <>
            <p>Name: {user.data.name}</p>
            <div>Identities: 
            {user.data.identities.map((item, index) => {
                return <span key={index}>{" "}
                    {item.provider}
                    {" "}
                </span>;
            })}
            </div>
        </>
    )
}
