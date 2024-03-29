import React, { useEffect, useState } from 'react';
import { Container, Typography, TextField, Button, IconButton } from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';

const EditableLabel = ({profile,onUpdate}) =>{

    const [editMode, setEditMode] = useState(false);
  
    const [labelValue, setLabelValue] = useState('');
    
    useEffect(()=>{
        setLabelValue(profile);
    },[profile])

    const handleTextChange = (event) => {
        setLabelValue(event.target.value);
        onUpdate(event.target.value);
    };

    const toggleEditMode = () => {
        setEditMode(!editMode);
    };

    const onClick = () => {
        console.log('hello');
    };

    return(
        <div style={{ display: 'flex', alignItems: 'center', gap: '10px',marginBottom: '16px' }}>
                {editMode ? (
                    <>
                    <TextField
                    value={labelValue}
                    type="password"
                    onChange={handleTextChange}
                    onBlur={toggleEditMode}
                    autoFocus
                    />
                    <Button onClick={onClick} variant="outlined">
                    Done
                </Button>
                </>
                ) : (
                    <>
                        <span>{labelValue}</span>
                    <IconButton onClick={toggleEditMode} aria-label="edit">
                        <EditIcon />
                    </IconButton>
                </>
                )}
            </div>
    )
}

export default EditableLabel;

