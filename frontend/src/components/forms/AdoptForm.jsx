import {useState} from 'react';

export default function AdoptionForm(){
	const [pets, setPets]= useState([{}]);
	const [vet, setVets]= useState([{vetInfo:""}]);
	const [person, setPerson] = useState([{person: ""}]);
	const [inputs, setInputs]=useState({});


	const handlePetChange =(index,e) =>{
		const newPets = [...pets];
		newPets[index] = {...newPets[index], [e.target.name]: e.target.value};
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

	const handleVetChange =(index,e) =>{
		const newVets = [...vet];
		newVets[index] = {...newVets[index], [e.target.name]: e.target.value};
		setVets (newVets);

	}

	const addVetField = () =>{
		setVets((prev) => [...prev,""]);
	};

	const removeVet = (index) => {
		const newVets = [...vet];
		newVets.splice (index,1)
		setVets (newVets);

	};

	const handlePeopleChange =(index,e) =>{
		const newPerson = [...person];
		newPerson[index] = {...newPerson[index], [e.target.name]: e.target.value};
		setPerson (newPerson);

	}
	const addPeopleField = () =>{
		setPerson((prev) => [...prev,""]);
	};

	const removePeople = (index) => {
		const newPeople = [...person];
		newPeople.splice (index,1)
		setPerson (newPeople);

	};

	const handleChange = (event)=>{
		const name = event.target.name;
		const inputValue = event.target.value;
		setInputs(prevState => ({...prevState, [name]: inputValue}));
	}
	// Enter the api call
	function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}

	return (

		<form action={handleSubmit} className="adoption-form">
			<label htmlFor="firstname" >First Name
				<input
					type="text"
					id="firstName"
					name="first_name"
					value={inputs.first_name || ""}
					onChange={handleChange}
					required/>
			</label>

			<label htmlFor='LastName'>Last Name
				<input
					type="text"
					id="LastName"
					name="last_name"
					value={inputs.last_name || ""}
					onChange={handleChange}
					required
				/>
			</label>

			<label htmlFor='PhoneNumber'>Phone Number
				<input
					type="tel"
					id="PhoneNumber"
					name="adopter_number"
					value={inputs.adopter_number || ""}
					pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}"
					onChange={handleChange}
					required
				/> </label>

			<label htmlFor='age'>Age
				<input
					type="text"
					id="age"
					name="adopter_age"
					value={inputs.adopter_age || ""}
					onChange={handleChange}
					required
				/> </label>

			<label htmlFor='emailaddress'>Email Address
				<input
					type="email"
					id="emailaddress"
					name="adopter_email"
					value={inputs.adopter_email || ""}
					onChange={handleChange}
					required
				/> </label>

			<label htmlFor='physicalAddress'>Physical Address
				<input
					type="textarea"
					id='physicalAddress'
					name="physical_address"
					value={inputs.physical_address || ""}
					onChange={handleChange}
					required
				/> </label>

			<label htmlFor='mailingAddress'>Mailing Address
				<input
					type="textarea"
					id ='mailingAddress'
					name="mailing_address"
					value={inputs.mailing_address || ""}
					onChange={handleChange}
					required
				/></label>

			<label htmlFor='houseDistance'>Do you live more than 100 miles from Shreveport Louisiana?
				<input
					type="text"
					name="house_distance"
					id='houseDistance'
					value={inputs.house_distance || ""}
					onChange={handleChange}
					required
				/> </label>


			<label htmlFor='allergies'>Have you (Or anyone in you home) ever experienced pet Allergies</label>
				<input
					type="radio"
					id="allergies"
					name="allergies"
					value="Yes"
					checked={inputs.allergies === 'Yes'}
					onChange={handleChange}
					required/>
			<label>Yes</label>


			<label>
				<input type="radio"
					   name="allergies"
					   id='allergies'
					   value="No"
					   checked={inputs.allergies === 'No'}
					   onChange={handleChange}
				/>
				No
			</label>

			<label>If yes, Above please describe
				<input
					type="textarea"
					id="allergyDescription"
					name="allergies_yes"
					value={inputs.allergies_yes || ""}
					onChange={handleChange}
				/>
			</label>

			<label>Have you adopted from the HSNWLA before?
				<input
					type="text"
					id="prevadoption"
					name="prev_adoption"
					value={inputs.prev_adoption || ""}
					onChange={handleChange}
					required
				/> </label>

			<label>
				Name of the dog you are interested in adopting -
				Applications on dogs not available for adoption will be rejected.
				<input
					type="text"
					name="interested_dog"
					value={inputs.interested_dog || ""}
					onChange={handleChange}
					required
				/> </label>

			<label htmlFor="first_adopt">Would this be your first animal?
			</label>
			<input
				type="radio"
				id="first_adopt"
				name="prev_adoption"
				value='Yes'
				checked={inputs.prev_adoption === 'Yes'}
				onChange={handleChange}
			/>
			<label> Yes</label>
			<input
				type="radio"
				id='first_adopt'
				name="first_adopt"
				value = 'No'
				checked={inputs.prev_adoption === 'No'}
				onChange={handleChange}
				/>
			<label> No</label>



			<label htmlFor="prevsurrender">Have you ever surrendered a pet to an animal shelter or rescue?
				<input
					type="radio"
					name="prev_surrender"
					id="prevsurrender"
					value="Yes"
					checked={inputs.prev_surrender === 'Yes'}
					onChange={handleChange}
				/>
				Yes
			</label>
			<label>
				<input
					type="radio"
					id="prevsurrender"
					name="prev_surrender"
					value="No"
					checked={inputs.prev_surrender === 'No'}
					onChange={handleChange}
				/>
				No
			</label>

			<label htmlFor="YesSurrender">If "Yes" above, please explain why.
				<input
					type="text"
					name="yes_surrender"
					value={inputs.yes_surrender || ""}
					onChange={handleChange}
				/>
			</label>



			<h3 htmlFor="currentDogs">Enter the dog you currently/previously owned in the last 10 years</h3>
			{pets.map((element, index = 0) => (
				<div key={index}>

					<label>Name of Pet
						<input
							type="text"
							name="pet_name"
							value={element.pet_name || ""}
							onChange={e => handlePetChange(index, e)}/>
					</label>



					<label>Pet Breed
						<input type="text"
							   name="pet_breed"
							   value={element.pet_breed || ""}
							   onChange={e => handlePetChange(index, e)}/>
					</label>


					<label>Timed Owned
						<input
							type="text"
							name="time_owned"
							value={element.time_owned || ""}
							onChange={e => handlePetChange(index, e)}

						/> </label>


					<label htmlFor='curr_owned'>Do you still own this pet?
						<input
							type="radio"
							id='curr_owned'
							name="still_owned"
							value="Yes"
							checked={element.still_owned === 'Yes'}
							onChange={e => handlePetChange(index, e)}
						/>
						Yes
					</label>

					<label htmlFor="curr_owned">
						<input
							type="radio"
							id='curr_owned'
							name="still_owned"
							value="No"
							checked={element.still_owned === 'No'}
							onChange={e => handlePetChange(index, e)}
						/>
						No
					</label>


					<label>Date of No Longer in Ownership of Passing of Pet
						<input
							type="date"
							name="no_longer_owned1_date"
							value={element.no_longer_owned1_date || ""}
							onChange={e => handlePetChange(index, e)}
							required
						/></label>

					<label htmlFor="whathappened">What Happened?
						<input
							type="textarea"
							id = 'whathappened'
							name="what_happened_to_owned_dog"
							value={element.what_happened_to_owned_dog || ""}
							onChange={e => handlePetChange(index, e)}
							required
						/></label>

					<label htmlFor='inside_outide'>Inside or Outside?
					<input
						type="radio"
						id = 'inside_outide'
						name="inside_outside_dog"
						value="inside"
						checked={element.inside_outside_dog === 'inside'}
						onChange={e => handlePetChange(index, e)}
						/>
					Inside
						<input
							type="radio"
							id = 'inside_outide'
							name="inside_outside_dog"
							value="outside"
							checked={element.inside_outside_dog === 'outside'}
							onChange={e => handlePetChange(index, e)}
						/>
					Outside
						</label>

					{pets.length-1=== index && pets.length<10 && <button className = "button add" type="button" onClick={()=> addPetField()}> Add Pet</button>}
					{index ? (
						<button type="button" className="button remove" onClick={() => removePet(index)}>
							Remove
						</button>
					) : null}

				</div>
			))}



			<label htmlFor='currentSpayedNeutered'>Are your current dogs/cats spayed and/or neutered?
			</label>
			<input
				type="radio"
				id='currentSpayedNeutered'
				name="current_spayed_neutered"
				value="Yes"
				checked={inputs.current_spayed_neutered === 'Yes'}
				onChange={handleChange}
			/>
			<label>Yes</label>
			<input
				type="radio"
				id='currentSpayedNeutered'
				name="current_spayed_neutered"
				value="No"
				checked={inputs.current_spayed_neutered === 'No'}
				onChange={handleChange}
			/>
			<label>No</label>
			<input
				type="radio"
				id='currentSpayedNeutered'
				name="current_spayed_neutered"
				value="N/A"
				checked={inputs.current_spayed_neutered === 'N/A'}
				onChange={handleChange}
			/>
			<label>N/A</label>



			<label>Do you purchase heartworm prevention from a vet?
			</label>
			<input
				type="radio"
				id="heartworm"
				name="heartworm_purchase"
				value="Yes"
				checked={inputs.heartworm_purchase === 'Yes'}
				onChange={handleChange}
			/>
			<label>Yes</label>
			<input
				type="radio"
				id='heartworm'
				name="heartworm_purchase"
				value="No"
				checked={inputs.heartworm_purchase === 'No'}
				onChange={handleChange}
			/>
			<label>No</label>



			<label htmlFor='noHeartwormResponse'>If you answered "No" above,
				please indicate where you purchase prevention from and
				provide the name brand.</label>
			<input
				type="textarea"
				id='noHeartwormResponse'
				name="no_heartworm"
				value={inputs.no_heartworm || ""}
				onChange={handleChange}
			/>




			<h3>Please Enter your Veterinarian's information</h3>
			{vet.map((element, index=1) => (
				<div key={index}>

					<label htmlFor='vetname'>Vet name
						<input
							type="text"
							id='vetname'
							name="vet_name"
							value={element.vet_name || ""}
							onChange={e => handleVetChange(index, e)}/>
					</label>

					<label htmlFor="vetNumber"> Phone Number
						<input
							type="tel"
							id='vetNumber'
							name="vet_number"
							value={element.vet_number || ""}
							   onChange={e => handleVetChange(index, e)} />
						</label>


						<label htmlFor='vetAssoiatedPet'>Pet associated with
							<input
								type="text"
								id='vetAssoiatedPet'
								name="pet_associated_vet"
								value={element.pet_associated_vet || ""}
								onChange={e => handleVetChange(index, e)}/>
						</label>
					{vet.length-1=== index && vet.length<10 && <button className = "button add" type="button" onClick={()=> addVetField()}> Add Vet</button>}
						{index ? (
							<button type="button" className="button remove" onClick={() => removeVet(index)}>
								Remove
							</button>
						) : null}
				</div>
				))}



			<h3> Please enter information for the people who live in your home</h3>
			{person.map((element,index =1)=>(
				<div key={index}>

					<label> Name
						<input
							type= "text"
							name="person_in_house_name"
							value={element.person_in_house_name || ""}
							onChange={e => handlePeopleChange(index, e)}/>
					</label>


					<label> Date of Birth
						<input type="date"
							   name ="DOB_person_house_name"
							   value={element.DOB_person_house_name || ""}
							   onChange={e => handlePeopleChange(index, e)}/>
					</label>

					<label>Relationship
						<input type="text"
							   name="relationship_person_house"
							   value={element.relationship_person_house|| ""}
							   onChange={e => handlePeopleChange(index, e)}/>
					</label>

					{person.length-1=== index && person.length<10 && <button className = "button add" type="button" onClick={()=> addPeopleField()}> Add Person</button>}
					{
						index ?
							<button type="button"  className="button-perosn-remove" onClick={() => removePeople(index)}>Remove</button>
							: null
					}
				</div>

			))}

			<label htmlFor='home-type'> What type of home do you live in?
					<input
						type="radio"
						id='home-type'
						name="type_of_house"
						value="House"
						checked={inputs.type_of_house === 'House'}
						onChange={handleChange}
						required/>
					House
				</label>
				<label>
					<input
						type="radio"
						id='home-type'
						name="type_of_house"
						value="Apartment"
						checked={inputs.type_of_house === 'Apartment'}
						onChange={handleChange}
					/>
					Apartment
				</label>
				<label>
					<input
						type="radio"
						id='home-type'
						name="type_of_house"
						value="Condo"
						checked={inputs.type_of_house === 'Condo'}
						onChange={handleChange}
					/>
					Condo
				</label>



			<label htmlFor='RentOrOwn'> Do you rent or Own?
					<input
						type="radio"
						id='RentOrOwn'
						name="rent_own"
						value="Rent"
						checked={inputs.rent_own === 'Rent'}
						onChange={handleChange}
						required/>
					Rent
				</label>

				<label>
					<input
						type="radio"
						id='RentOrOwn'
						name="rent_own"
						value="Own"
						checked={inputs.rent_own === 'Own'}
						onChange={handleChange}
					/>
					Own
				</label>



				<label> If you rent, list landlord's name and phone number
					<input
						type="text"
						name="landlord_name"
						value={inputs.landlord_name || ""}
						onChange={handleChange}

					/>
				</label>



				<label htmlFor='consent'>Yes, I give consent to the HSNWLA to contact your veterinarian(s) for records. (Initial Here)
					<input
						type="text"
						id="consent"
						name="HSNWLA_consent"
						value={inputs.HSNWLA_consent || ""}
						onChange={handleChange}
						required/></label>



				<label> Yes, prior to submitting this application, I have contacted
					my veterinarian(s) to give them authorization to release records to
					us (Initial Here)
					<input
						type="text"
						name="vet_authorization"
						value={inputs.vet_authorization || ""}
						onChange={handleChange}
						required

					/></label>



				<label htmlFor='personhome'> Is someone home during the day?
					<input
						type="radio"
						id='personhome'
						name="person_at_home"
						value="Yes"
						checked={inputs.person_at_home === 'Yes'}
						onChange={handleChange}
						required/>
					Yes
				</label>
				<label>
					<input
						type="radio"
						id='personhome'
						name="person_at_home"
						value="No"
						checked={inputs.person_at_home === 'No'}
						onChange={handleChange}
					/>
					No
				</label>



				<label> If so, who and when?
					<input
						type="text"
						name='person_at_home_yes'
						value={inputs.person_at_home_yes || ""}
						onChange={handleChange}
						required

					/>
				</label>



				<label htmlFor='fencedyard'> Do you have a fenced yard?
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="Yes"
						checked={inputs.fenced_yard === 'Yes'}
						onChange={handleChange}/>
					Yes, Chain Link
				</label>

				<label>
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="Yes, Privacy"
						checked={inputs.fenced_yard === 'Yes, Privacy'}
						onChange={handleChange}
					/>
					Yes, Privacy
				</label>

				<label>
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="Yes, Wireless Electric"
						checked={inputs.fenced_yard === 'Yes, Wireless Electric'}
						onChange={handleChange}
					/>
					Yes, Wireless Electric
				</label>

				<label>
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="Yes, Other"
						checked={inputs.fenced_yard === 'Yes, Other'}
						onChange={handleChange}
					/>
					Yes, Other
				</label>

				<label>
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="Yes, Partial"
						checked={inputs.fenced_yard === 'Yes, Partial'}
						onChange={handleChange}
					/>
					Yes, Partial
				</label>

				<label>
					<input
						type="radio"
						id='fencedyard'
						name="fenced_yard"
						value="No"
						checked={inputs.fenced_yard === 'No'}
						onChange={handleChange}
					/>
					No
				</label>



				<label htmlFor='dogstay'>Where will the dog stay when you are gone for the day?
					<input
						type="radio"
						id='dogstay'
						name="dog_stay_while_away"
						value="Inside"
						checked={inputs.dog_stay_while_away === 'Inside'}
						onChange={handleChange}
					/>
					Inside
				</label>

				<label htmlFor='dogstay'>
					<input
						type="radio"
						id='dogstay'
						name="dog_stay_while_away"
						value="Outside"
						checked={inputs.dog_stay_while_away === 'Outside'}
						onChange={handleChange}
					/>
					Outside
				</label>

				<label htmlFor="dogStayAtNight">Where will the dogs sleep at night?
					<input
						type="radio"
						id="dogStayAtNight"
						name="where_will_sleep"
						value="Inside"
						checked={inputs.where_will_sleep === 'Inside'}
						onChange={handleChange}
						/>
					Inside
				</label>

				<label htmlFor="dogStayAtNight">
					<input
						type="radio"
						id='dogStayAtNight'
						name="where_will_sleep"
						value="Outside"
						checked={inputs.where_will_sleep === 'Outside'}
						onChange={handleChange}
					/>
					Outside
				</label>

				<label> What are your thoughts on a dog living outdoors?
					<input type="textarea"
						   name="thoughts_on_inside_outside"
						   value={inputs.thoughts_on_inside_outside || ""}
						   onChange={handleChange}

					/></label>


				<label> Where do dogs get heartworm disease from?
					<input
						type="textarea"
						name='heartworm_knowledge'
						value={inputs.heartworm_knowledge || ""}
						onChange={handleChange}
					/></label>

				<label htmlFor='dogsHeartwormProtected'> Are your dogs current on heartworm prevention?
					<input
						type="radio"
						id='dogsHeartwormProtected'
						name="heartworm_prevention"
						value='Yes'
						checked={inputs.heartworm_prevention === 'Yes'}
						onChange={handleChange}
						required/>
					Yes
				</label>

				<label>
					<input
						type="radio"
						id='dogsHeartwormProtected'
						name="heartworm_prevention"
						value="No"
						checked={inputs.heartworm_prevention === 'No'}
						onChange={handleChange}
					/>
					No
				</label>

				<label>
					<input
						type="radio"
						id='dogsHeartwormProtected'
						name="heartworm_prevention"
						value="N/A"
						checked={inputs.heartworm_prevention === 'N/A'}
						onChange={handleChange}
					/>
					N/A
				</label>


				<label htmlFor='activeMilitary'> Are you active military?
					<input
						type="radio"
						id='activeMilitary'
						name="active_duty"
						value='Yes'
						checked={inputs.active_duty === 'Yes'}
						onChange={handleChange}
						required/>
					Yes
				</label>

				<label>
					<input
						type="radio"
						id='activeMilitary'
						name="active_duty"
						value='No'
						checked={inputs.active_duty === 'No'}
						onChange={handleChange}
					/>
					No
				</label>


				<label htmlFor='ifActiveDuty'> If active duty, is there a chance of deployment?
					<input
						type="radio"
						id='ifActiveDuty'
						name="chance_of_deployment"
						value='Yes'
						checked={inputs.chance_of_deployment === 'Yes'}
						onChange={handleChange}
					/>
					Yes
				</label>

				<label>
					<input
						type="radio"
						id='ifActiveDuty'
						name="chance_of_deployment"
						value="No"
						checked={inputs.chance_of_deployment === 'No'}
						onChange={handleChange}
					/>
					No
				</label>


				<label>If you are active duty military and deploy, where will your pets go?
					<input
						type="text"
						name="where_will_pets_go"
						value={inputs.where_will_pets_go || ""}
						onChange={handleChange}
					/> </label>


				<label>
					INCOMPLETE APPLICATIONS WITH MISSING INFORMATION WILL
					BE RETURNED AS INCOMPLETE AND NOT PROCESSED. PLEASE MAKE SURE
					ALL FIELDS ARE ACCURATELY FILLED OUT AND YOUR VETERINARIAN HAS
					BEEN CONTACTED TO RELEASE RECORDS.
					<input
						type="checkbox"
						name="incomplete_form"
						value={inputs.incomplete_form || ""}
						onChange={handleChange}
						required/>
					Yes I, agree
				</label>


				<label>
					I UNDERSTAND AND AM PREPARED FOR THE FINANCIAL RESPONSIBILITY
					OF OWNING A PET INCLUDING, BUT NOT LIMITED TO ANNUAL VACCINATIONS, AND
					HEARTWORM PREVENTION.
					<input
						type="checkbox"
						name="financial_responsibility"
						value={inputs.financial_responsibility || ""}
						onChange={handleChange}
						required/>
					Yes I, agree
				</label>


				<label>PLEASE NOTE THAT IF YOU SHOULD BECOME UNABLE TO CARE FOR THIS ANIMAL,
					IT MUST BE RETURNED TO THE HSNWLA. IT CAN NOT BE GIVEN AWAY TO ANYONE.
					(Initial Below)
					<input
						type="checkbox"
						name='return_options'
						value={inputs.return_options || ""}
						onChange={handleChange}
						required/>
					Yes I, agree
				</label>



				<label> Signature</label>
				<input type="text"
					   name='signature'
					   value={inputs.signature || ""}
					   onChange={handleChange}
					   required />


				<button type="submit">Submit</button>

		</form>
)
}
