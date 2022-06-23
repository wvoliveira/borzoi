import {Typography} from "@mui/material";
import useUser from "../lib/utils/useUser";
    import EmailIcon from '@mui/icons-material/Email';

export default function ProfileInfo() {
    const { user } = useUser({
        redirectTo: "/login",
    });

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
            {user.data.identities.map(item => {
                return <>{" "}
                    {item.provider}
                    {" "}
                </>;
            })}
            </div>
        </>
    )
}
