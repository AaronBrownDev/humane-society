import {useState} from 'react';

export default function AdoptionForm(){
	const [pets, setPets]= useState([{currentDogs: " "}]);
	const [vet, setVets]= useState([{vetInfo:""}]);
	const [person, setPerson] = useState([{person: ""}]);

	const handlePetChange =(index,value) =>{
		const newPets = [...pets];
		newPets [index] = value;
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

	const handleVetChange =(index,value) =>{
		const newVets = [...pets];
		newVets [index] = value;
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

	const handlePeopleChange =(index,value) =>{
		const newPerson = [...person];
		newPerson [index] = value;
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
	
	function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}

	return (

		<form action={handleSubmit} className="adoption-form">
			<label>Date <input type="date" id="date" name="date" required/></label>

			<label>First Name<input type="text" id="firstName" name="first_name" required/> </label>

			<label>Last Name<input type="text" id="LastName" name="last_name"/> </label>

			<label>Phone Number <input type="text" id="PhoneNumber" name="adopter_number"/> </label>

			<label>Age<input type="number" id="age" name="adopter_age"/> </label>

			<label>Email Address<input type="email" id="emailaddress" name="adopter_email"/> </label>

			<label>Physical Address<input type="textarea" name="physical_address"/> </label>

			<label>Mailing Address<input type="textarea" name="mailing_address"/></label>

			<label>Do you live more than 100 miles from Shreveport Louisiana?
				<input type="text" name="house_distance"/> </label>


			<label>Have you (Or anyone in you home) ever experienced pet Allergies
				<input type="radio" id="allergies" name="allergies" value="Yes" required/>
				Yes
			</label>

			<label>
				<input type="radio" name="allergies" value="No"/>
				No
			</label>

			<label>If yes, Above please describe <input type="textarea" id="allergyDescription" name="allergies_yes"/>
			</label>

			<label>Have you adopted from the HSNWLA before?
				<input type="text" id="prevadoption" name="prev_adoption"/> </label>

			<label>
				Name of the dog you are interested in adopting -
				Applications on dogs not available for adoption will be rejected.
				<input type="text" name="interested_dog"/> </label>

			<label for="prevadoption">Would this be your first animal?
			</label>
			<input type="radio" id="first_adopt" name="prev_adoption"/>
			<label> Yes</label>
			<input type="radio" name="first_adopt"/>
			<label> No</label>



			<label for="prevsurrender">Have you ever surrendered a pet to an animal shelter or rescue?
				<input type="radio" name="prev_surrender" value="Yes"/>
				Yes
			</label>
			<label>
				<input type="radio" name="prev_surrender" value="No"/>
				No
			</label>

			<label for="YesSurrender">If "Yes" above, please explain why.<input type="text" name="yes_surrender"/>
			</label>



			<h3 htmlFor="currentDogs">Enter the dog you currently/previously owned in the last 10 years</h3>
			{pets.map((element, index = 1) => (
				<div key={index}>

					<label>Name of Pet
						<input type="text" name="pet_name" value={element.pet_name || ""}
							   onChange={e => handlePetChange(index, e)}/>
					</label>



					<label>Pet Breed
						<input type="text" name="pet_breed" value={element.pet_breed || ""}
							   onChange={e => handlePetChange(index, e)}/>
					</label>


					<label>Timed Owned <input type="text" name="time_owned"/> </label>


					<label>Do you still own this pet?
						<input type="radio" name="still_owned" value="Yes"/>
						Yes
					</label>

					<label htmlFor="No">
						<input type="radio" name="still_owned" value="No"/>
						No
					</label>


					<label>Date of No Longer in Ownership of Passing of Pet
						<input type="date" name="no_longer_owned1_date"/></label>

					<label htmlFor="whathappened">What Happened?
						<input type="textarea" name="what_happened_to_owned_dog1"/></label>

					<label>Inside or Outside? </label>
					<input type="radio" name="inside_outside_dog1" value="inside"/>
					<label>Inside</label>
					<input type="radio" name="inside_outside_dog1" value="outside"/>
					<label>Outside</label>

					{pets.length-1=== index && pets.length<10 && <button className = "button add" type="button" onClick={()=> addPetField()}> Add</button>}
					{index ? (
						<button type="button" className="button remove" onClick={() => removePet(index)}>
							Remove
						</button>
					) : null}

				</div>
			))}



			<label>Are your current dogs/cats spayed and/or neutered?
			</label>
			<input type="radio" name="current_spayed_neutered" value="Yes"/>
			<label>Yes</label>
			<input type="radio" name="current_spayed_neutered" value="No"/>
			<label>No</label>
			<input type="radio" name="current_spayed_neutered" value="N/A"/>
			<label>N/A</label>



			<label>Do you purchase heartworm prevention from a vet?
			</label>
			<input type="radio" name="heartworm_purchase" value="Yes"/>
			<label>Yes</label>
			<input type="radio" name="heartworm_purchase" value="No"/>
			<label>No</label>



			<label>If you answered "No" above,
				please indicate where you purchase prevention from and
				provide the name brand.</label>
			<input type="textarea" name="no_heartworm"/>




			<h3>Please Enter your Veterinarian's information</h3>
			{vet.map((element, index=1) => (
				<div key={index}>

					<label>Vet name
						<input type="text" name="vet_name" value={element.vet_name || ""}
							   onChange={e => handleVetChange(index, e)}/>
					</label>

					<label> Phone Number
						<input type="text" name="vet_number" value={element.vet_number || ""}
							   onChange={e => handleVetChange(index, e)} />
						</label>


						<label>Pet associated with
							<input type="text" name="pet_associated_vet"
								   value={element.pet_associated_vet || ""}
								   onChange={e => handleVetChange(index, e)}/>
						</label>
					{vet.length-1=== index && vet.length<10 && <button className = "button add" type="button" onClick={()=> addVetField()}> Add</button>}
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
						<input type= "text" name="person-in-house-name" value={element.pet_associated_vet || ""}
							   onChange={e => handlePeopleChange(index, e)}/>
					</label>


					<label> Date of Birth
						<input type="date" name ="DOB-person-house-name" value={element.pet_associated_vet || ""}
							   onChange={e => handlePeopleChange(index, e)}/>
					</label>

					<label>Relationship
						<input type="text" name="relationship-person-house" value={element.pet_associated_vet || ""}
							   onChange={e => handlePeopleChange(index, e)}/>
					</label>

					{person.length-1=== index && person.length<10 && <button className = "button add" type="button" onClick={()=> addPeopleField()}> Add</button>}
					{
						index ?
							<button type="button"  className="button-perosn-remove" onClick={() => removePeople(index)}>Remove</button>
							: null
					}
				</div>

			))}






			<label> What type of home do you live in?
					<input type="radio" name="type_of_house" value="House" required/>
					House
				</label>
				<label>
					<input type="radio" name="type_of_house" value="Apartment"/>
					Apartment
				</label>
				<label>
					<input type="radio" name="type_of_house" value="Condo"/>
					Condo
				</label>



			<label> Do you rent or Own?
					<input type="radio" name="rent_own" value="Rent" required/>
					Rent
				</label>

				<label>
					<input type="radio" name="rent_own" value="Own"/>
					Own
				</label>



				<label> If you rent, list landlord's name and phone number<input type="text"/></label>



				<label>Yes, I give consent to the HSNWLA to contact your veterinarian(s) for records. (Initial Here)
					<input type="text" name="HSNWLA_consent" required/></label>



				<label> Yes, prior to submitting this application, I have contacted
					my veterinarian(s) to give them authorization to release records to
					us (Initial Here)
					<input type="text"/></label>



				<label> Is someone home during the day?
					<input type="radio" name="person_at_home" value="Yes" required/>
					Yes
				</label>
				<label>
					<input type="radio" name="person_at_home" value="No"/>
					No
				</label>



				<label> If so, who and when?
					<input type="text"/>
				</label>



				<label> Do you have a fenced yard?
					<input type="radio" name="fenced_yard" value="Yes, Chain Link" required/>
					Yes, Chain Link
				</label>

				<label>
					<input type="radio" name="fenced_yard" value="Yes, Privacy"/>
					Yes, Privacy
				</label>

				<label>
					<input type="radio" name="fenced_yard" value="Yes, Wireless Electric"/>
					Yes, Wireless Electric
				</label>

				<label>
					<input type="radio" name="fenced_yard" value="Yes, Other"/>
					Yes, Other
				</label>

				<label>
					<input type="radio" name="fenced_yard" value="Yes, Partial"/>
					Yes, Partial
				</label>

				<label>
					<input type="radio" name="fenced_yard" value="No"/>
					No
				</label>



				<label>Where will the dog stay when you are gone for the day?
					<input type="radio" name="dog_stay_while_away" value="Inside" required/>
					Inside
				</label>

				<label>
					<input type="radio" name="dog_stay_while_away" value="Outside"/>
					Outside
				</label>

				<label>Where will the dogs sleep at night?
					<input type="radio" name="where_will_sleep" value=" Inside" required/>
					Inside
				</label>

				<label>
					<input type="radio" name="where_will_sleep" value=" Outside"/>
					Outside
				</label>

				<label> What are your thoughts on a dog living outdoors?<input type="textarea"
																			   name="thoughts_on_inside_outside"/></label>


				<label> Where do dogs get heartworm disease from? <input type="textarea"/></label>

				<label> Are your dogs current on heartworm prevention?
					<input type="radio" name="heartworm_prevention" required/>
					Yes
				</label>

				<label>
					<input type="radio" name="heartworm_prevention"/>
					No
				</label>

				<label>
					<input type="radio" name="heartworm_prevention"/>
					N/A
				</label>


				<label>
					Is there anything else you would like us to know about your home?
					<input type="textarea"/>
				</label>


				<label> Are you active military?
					<input type="radio" name="active_duty" required/>
					Yes
				</label>

				<label>
					<input type="radio" name="active_duty"/>
					No
				</label>


				<label> If active duty, is there a chance of deployment?
					<input type="radio" name="chance_of_deployment"/>
					Yes
				</label>

				<label>
					<input type="radio" name="chance_of_deployment"/>
					No
				</label>


				<label>If you are active duty military and deploy, where will your pets go? <input
					type="text"/> </label>


				<label>
					INCOMPLETE APPLICATIONS WITH MISSING INFORMATION WILL
					BE RETURNED AS INCOMPLETE AND NOT PROCESSED. PLEASE MAKE SURE
					ALL FIELDS ARE ACCURATELY FILLED OUT AND YOUR VETERINARIAN HAS
					BEEN CONTACTED TO RELEASE RECORDS.
					<input type="checkbox" required/>
					Yes I, agree
				</label>


				<label>
					I UNDERSTAND AND AM PREPARED FOR THE FINANCIAL RESPONSIBILITY
					OF OWNING A PET INCLUDING, BUT NOT LIMITED TO ANNUAL VACCINATIONS, AND
					HEARTWORM PREVENTION.
					<input type="checkbox" required/>
					Yes I, agree
				</label>


				<label>PLEASE NOTE THAT IF YOU SHOULD BECOME UNABLE TO CARE FOR THIS ANIMAL,
					IT MUST BE RETURNED TO THE HSNWLA. IT CAN NOT BE GIVEN AWAY TO ANYONE.
					(Initial Below)
					<input type="checkbox" required/>
					Yes I, agree
				</label>



				<label> Signature</label>
				<input type="text" required/>


				<button type="submit">Submit</button>

		</form>
)
}
