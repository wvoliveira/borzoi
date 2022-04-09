import {Typography} from "@mui/material";
import useUser from "../lib/utils/useUser";
    import EmailIcon from '@mui/icons-material/Email';

export default function ProfileInfo() {
    const { user } = useUser({
        redirectTo: "/login",
    });

    if (!user) {
        return (
            <Typography>
                Loading...
            </Typography>
        )
    }

    return (
        <Typography>
            Name: {user.data.name}<br/>
            Identities:<br/>
            {user.data.identities.map(item => {
                if (item.provider === "email") {
                    return <EmailIcon key={item.provider} />;
                    { " " }
                }
                return <></>;
            })}
        </Typography>
    )
}
