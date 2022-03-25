import * as React from 'react';
import {DataGrid, GridActionsCellItem} from '@mui/x-data-grid';
import DeleteIcon from '@mui/icons-material/Delete';
import LinkIcon from '@mui/icons-material/Link';
import EditIcon from '@mui/icons-material/Edit';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import {Typography} from "@mui/material";
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Stack from '@mui/material/Stack';
import Skeleton from '@mui/material/Skeleton';
import {useSWRConfig} from "swr";

function LinkList({props, error}) {
    const {mutate} = useSWRConfig();

    const [loading, setLoading] = React.useState(false);
    const [openDeleteDialog, setOpenDeleteDialog] = React.useState({show: false, id: ""});

    function shortenerLink(params) {
        let url;
        if (params.row.domain.startsWith("localhost")) {
            url = new URL("http://" + params.row.domain + "/" + params.row.keyword);
            return url;
        }
        url = new URL("https://" + params.row.domain + "/" + params.row.keyword);
        return url;
    }

    const openLink = React.useCallback(
        (params) => () => {
            let url = shortenerLink(params);
            window.open(url, '_blank').focus();
        },
        [],
    );

    const copyLink = React.useCallback(
        (e) => () => {
            let url = shortenerLink(e);
            navigator.clipboard.writeText(url);
        },
        [],
    )

    const editLink = React.useCallback(
        (id) => () => {
            console.log("Edit link function");
        },
        [],
    );

    const deleteLink = React.useCallback(
        (id) => () => {
            setOpenDeleteDialog({show: true, id: id});
        },
        [],
    );

    const handleCancel = (e) => {
        setOpenDeleteDialog({show: false});
    };

    const handleDelete = (e) => {
        setLoading(true);

        fetch("/api/v1/" + openDeleteDialog.id, {
            method: 'DELETE',
        })
            .then(response => {
                setLoading(false);
                return response.json();
            })
            .then(response => {
                if (response.status === 'successful') {
                    console.log('request OK!');
                } else {
                    let message = 'status: ' + response.status + ' message: ' + response.message
                    console.log(message);
                }
                mutate('/api/v1/links')
            });

        setOpenDeleteDialog({show: false});
    };

    const columns = [
        {field: 'id', headerName: 'ID', width: 180},
        {field: 'domain', headerName: 'Domain', width: 120},
        {field: 'keyword', headerName: 'Keyword', width: 100},
        {field: 'url', headerName: 'URL', flex: 1},
        {field: 'title', headerName: 'Title', flex: 1},
        {
            field: 'created_at',
            headerName: 'Created',
            type: 'dateTime',
            width: 200,
            valueGetter: ({value}) => value && new Date(value)
        },
        {
            field: 'updated_at',
            headerName: 'Updated',
            type: 'dateTime',
            width: 200,
            valueGetter: ({value}) => value && new Date(value)
        },
        {field: 'active', headerName: 'Active', type: 'boolean', width: 100},
        {
            field: 'actions',
            headerName: 'Actions',
            type: 'actions',
            minWidth: 80,
            getActions: (params) => [
                <GridActionsCellItem
                    icon={<LinkIcon/>}
                    label="Open link"
                    onClick={openLink(params)}
                    title={"Open link"}

                />,
                <GridActionsCellItem
                    icon={<ContentCopyIcon/>}
                    label="Copy link"
                    onClick={copyLink(params)}
                    title={"Copy link"}
                />,
                <GridActionsCellItem
                    icon={<EditIcon/>}
                    label="Edit"
                    onClick={editLink(params)}
                    title={"Edit"}
                    showInMenu={true}
                />,
                <GridActionsCellItem
                    icon={<DeleteIcon/>}
                    label="Delete"
                    onClick={deleteLink(params.id)}
                    showInMenu={true}
                />,
            ],
        },
    ];

    if (!props) {
        return (
            <Stack spacing={2}>
                <Skeleton variant="rectangular" animation="wave" width='100%' height={70}/>
                {Array.from(new Array(7)).map((value, index) => (
                    <Skeleton key={index} variant="rectangular" animation="wave" width='100%' height={30}/>
                ))}
            </Stack>
        )
    }

    if (error) {
        console.log(error);
        return <Typography variant="subtitle1" component="subtitle1">
            Error to load: {error}
        </Typography>;
    }

    if (props) {
        console.log(props);
    }

    return (
        <div style={{height: 500, width: '100%'}}>
            <DataGrid
                loading={loading}
                rows={props.data.links == null ? [] : props.data.links}
                columns={columns}
                pageSize={10}
                rowsPerPageOptions={[10]}
                checkboxSelection={false}
            />
            <Dialog
                open={openDeleteDialog.show}
                onClose={handleCancel}
            >
                <DialogTitle>
                    {"Are you sure?"}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        The link will not work anymore and metrics will be deleted.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCancel}>Cancel</Button>
                    <Button onClick={handleDelete} autoFocus>Delete</Button>
                </DialogActions>
            </Dialog>
        </div>

    );
}

export default LinkList;
