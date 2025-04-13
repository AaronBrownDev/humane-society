import React from 'react';


export default function AddRemovePet() {
    const [pets, setPets]= React.useState([""]);
    const handlePetChange =(index,value) =>{
		const newPets = [...pets];
		newPets [index = value];
		setPets (newPets);
		
	}
	
	const addPetField = () =>{
	setPets((prev) => [...prev,""]);
	};
	
	const removePet = (index) => {
		const newPets = [...pets];
		newPets.splice (index,1)
		setPets (newPets);
		
		};
    
    return {
        
    }
}
