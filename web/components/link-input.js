import * as React from 'react';
import {useSWRConfig} from 'swr'
import {FormControl, InputLabel} from '@mui/material';
import {LoadingButton} from '@mui/lab';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import TextField from '@mui/material/TextField';
import LinearProgress from '@mui/material/LinearProgress';


function LinkInput({props, error}) {
    const {mutate} = useSWRConfig();

    const [domain, setDomain] = React.useState('');
    const [keyword, setKeyword] = React.useState('');
    const [url, setURL] = React.useState('');
    const [title, setTitle] = React.useState('');

    const [loadingCreatingLink, setLoadingCreatingLink] = React.useState(false);

    const handleChangeDomain = (e) => {
        setDomain(e.target.value);
    };

    const handleChangeKeyword = (e) => {
        setKeyword(e.target.value);
    };

    const handleChangeURL = (e) => {
        setURL(e.target.value);
    };

    const handleChangeTitle = (e) => {
        setTitle(e.target.value);
    };

    function clearInputs() {
        setDomain('');
        setKeyword("");
        setURL("");
        setTitle("");
    }

    const handleCreateLink = async (event) => {
        setLoadingCreatingLink(true);

        const data = {
            "domain": domain,
            "keyword": keyword,
            "url": url,
            "title": title
        }

        fetch("/api/v1/links", {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })
            .then(response => {
                setLoadingCreatingLink(false);
                return response.json();
            })
            .then(response => {
                if (response.status === 'successful') {
                    console.log('request OK!');
                } else {
                    let message = 'status: ' + response.status + ' message: ' + response.message
                    console.log(message);
                }
            })
            .then(() => {
                clearInputs();
                mutate('/api/v1/links')
            });

    };

    return (
        <div>
            <form onSubmit={handleCreateLink}>
                {props ? <FormControl required sx={{mr: 1, mb: 1, minWidth: 120}} size="small">
                        <InputLabel>Domain</InputLabel>
                        <Select
                            labelId="select-domain"
                            id="domain"
                            value={domain}
                            label="Domain"
                            onChange={handleChangeDomain}
                        >
                            {props.data.domains.map((item) => {
                                return <MenuItem key={item} value={item}>{item}</MenuItem>
                            })}
                        </Select>
                    </FormControl>
                    :
                    <FormControl required sx={{mr: 1, mb: 1, minWidth: 120}} size="small">
                        <InputLabel>Domain<LinearProgress/></InputLabel>
                        <Select
                            labelId="select-domain"
                            id="domain"
                            value='...'
                            label="Domain"
                            onChange={handleChangeDomain}
                        >
                            <MenuItem key="..." value='...'>{}</MenuItem>
                            <LinearProgress/>
                        </Select>
                    </FormControl>
                }
                <FormControl required sx={{mr: 1, mb: 1, minWidth: 120}}>
                    <TextField id="keyword" label="Keyword" variant="outlined" size="small" value={keyword}
                               onChange={handleChangeKeyword}/>
                </FormControl>
                <FormControl required sx={{mr: 1, mb: 1, minWidth: 120}}>
                    <TextField id="url" label="URL to redirect" variant="outlined" size="small" value={url}
                               onChange={handleChangeURL}/>
                </FormControl>
                <FormControl required sx={{mr: 1, mb: 1, minWidth: 300}}>
                    <TextField id="title" label="Title" variant="outlined" size="small" value={title}
                               onChange={handleChangeTitle}/>
                </FormControl>
                <FormControl required sx={{mr: 1, mb: 1, minWidth: 50}} size="small">
                    <LoadingButton
                        variant="contained"
                        onClick={handleCreateLink}
                        loading={loadingCreatingLink}
                    >Create</LoadingButton>
                </FormControl>
            </form>
        </div>
    );
}

export default LinkInput;
