
export default function adoptionForm(formData){
	
	/*const firstName = formData.get("first_name")
	const lastName = formData.get("last_name")
	const number = formData.get ("adopter_number")
	const age = formData.get ("adopter_age")
	const email = formData.get ("adopter_email")
	const physicalAddress = formData.get ("physical_address")
	const mailingAddress = formData.get ("mailing_address")
	const houseDistance = formData.get ("house_distance")
	const adoptersAllergies = formData.get ("allergies")
	const yesAdopterAllergies = formData.get ("allergies_yes")
	const previousAdoption = formData.get("prev_adoption")
	const interestedDog = formData.get("interested_Dog")
	const firstAdoption = formData.get ("first_adopt")
	const prevsurrender= formData.get ("prev_surrender")
	const yesSurrender = formData.get("yes_surrender")
	
	const data = Object.fromEntries(formData)
	console.log(data)
	
	*/
	
	
	return (

		<form action={adoptionForm}>
						<label>Date <input type= "date" id="date" name="date" required/></label><br></br>
						
						<label>First Name<input type="text" id= "firstName" name="first_name" required/> </label><br></br>

						<label>Last Name<input type="text" id= "LastName" name= "last_name"/> </label><br></br>
						
						<label>Phone Number <input type="number" id= "PhoneNumber" name="adopter_number"/> </label><br></br>
						
						<label>Age<input type="number" id= "age" name="adopter_age"/> </label><br></br>

						<label>Email Address<input type="email" id= "emailaddress" name ="adopter_email"/> </label><br></br>

						<label >Physical Address<input type="text"name = "physical_address"/> </label><br></br>
						
						<label>Mailing Address<input type="textarea" name="mailing_address"/></label><br></br>
						
						<label>Do you live more than 100 miles from Shreveport Louisiana?  
						<input type="text" name="house_distance"/> </label><br></br>
						

						<label for="allergies">Have you (Or anyone in you home) ever experienced pet Allergies
							<input type="radio" id= "allergies" name="allergies" value="Yes" required/> 
							Yes 
						</label>
						
						<label>
							<input type="radio" name="allergies" value="No"/>
							No
						</label><br></br>

						<label>If yes, Above please describe <input type="textarea" id= "allergyDescription" name="allergies_yes"/> </label><br></br>

						<label for="prevadoption">Have you adopted from the HSNWLA before?
						<input type="text" id= "prevadoption" name="prev_adoption"/> </label><br></br>
						
						<label for="interestedDog">
							Name of the dog you are interested in adopting - 
							Applications on dogs not available for adoption will be rejected.
						<input type="text" name= "interested_dog"/> </label><br></br>

						<label for="prevadoption">Would this be your first animal?
						</label>
						<input type="radio" id= "first_adopt" name="prev_adoption"/> 
						<label> Yes</label>
						<input type="radio" name="first_adopt"/>
						<label> No</label>
						<br></br>
						<br></br>

						<label for="prevsurrender">Have you ever surrendered a pet to an animal shelter or rescue?
							<input type="radio" name= "prev_surrender" value ="Yes"/> 
							Yes
						</label>
						<label>
							<input type="radio" name= "prev_surrender" value ="No"/> 
							No
						</label>
						<br></br>
						<label for="YesSurrender">If "Yes" above, please explain why.<input type="text" name= "yes_surrender"/> </label>
						<br></br>
						
						
						// Try the email component thing 
						<h1> Information on Pet 1</h1>
						<label >Name of Pet<input type="text" name="pet1_name"/></label><br></br>
						
						<label >Pet Breed <input type="text" name="pet1_breed"/> </label><br></br>
						
						
						<label>Timed Owned <input type="text" name="time_owned"/> </label><br></br>
						
						<label>Do you still own this pet? 
							<input type="radio" name="still_owned" value = "Yes"/>
							Yes
						</label>
						
						<label for ="No">
							<input type="radio" name="still_owned" value ="No"/>
							Yes
						</label><br></br>
					

						<label>Date of No Longer in Ownership of Passing of Pet 
						<input type="date" name="no_longer_owned1_date"/></label><br></br>
						
						<label for ="whathappened">What Happened? 
						<input type="textarea" name= "what_happened_to_owned_dog1"/></label><br></br>
						

						<label >Inside or Outside? </label>
						<input type="radio" name="inside_outside_dog1" value="inside"/> 
						<label>Inside</label>
						<input type="radio" name="inside_outside_dog1" value="outside"/> 
						<label>Outside</label>
						<br></br>
						<br></br>

						
						<h1> Information on Pet 2</h1>
						<label >Pet Breed </label>
						<input type="text" name="pet2_breed"/> 
						<br></br>
						<br></br>
						
						<label>Timed Owned 
						</label>
						<input type="text" name="time_owned"/> 
						<br></br>
						<br></br>
						
						<label for="prevadoption">Inside or Outside? 
						</label>
						<input type="radio" name="inside_outside_dog2" value="insde"/> 
						<label>Inside</label>
						<input type="radio" name="inside_outside_dog2" value="outside"/> 
						<label>Outside</label>
						<br></br>
						<br></br>
						
						<label for ="stillowned">Do you still own this pet? </label>
						<input type="radio" name="stilled_yes_no" value= "Yes"/>
						<label for ="Yes">Yes</label>
						<input type="radio" name="stilled_yes_no" value= "No"/>
						<label for ="No">Yes</label>
						<br></br>
						<br></br>
					

						<label>Date of No Longer in Ownership of Passing of Pet </label>
						<input type="date" name="no_longer_owned2_date"/>
						<br></br>
						<br></br>
						
						<label >What Happened? </label>
						<input type="text" name= "what_happened_to_owned_dog2"/>
						<br></br>
						<br></br>
						

					

						<h1> Information on Pet 3</h1>
						
						<label >Pet Breed </label>
						<input type="text" name="pet3_breed"/> 
						<br></br>
						<br></br>
						
						<label>Timed Owned 
						</label>
						<input type="text" name="time_owned"/> 
						<br></br>
						<br></br>
						
						<label>Inside or Outside? 
						</label>
						<input type="radio" name="inside_outside_dog3" value="insde"/> 
						<label>Inside</label>
						<input type="radio" name="inside_outside_dog3" value="outside"/> 
						<label>Outside</label>
						<br></br>
						<br></br>
						
						<label for ="stillowned">Do you still own this pet? </label>
						<input type="radio" name="still_yes_no3" value= "Yes"/>
						<label for ="Yes">Yes</label>
						<input type="radio" name="still_yes_no3" value= "No"/>
						<label for ="No">Yes</label>
						<br></br>
						<br></br>
					

						<label>Date of No Longer in Ownership of Passing of Pet </label>
						<input type="date" name="no_longer_owned3_date"/>
						<br></br>
						<br></br>
						
						<label >What Happened? </label>
						<input type="text" name= "what_happened_to_owned_dog3"/>
						<br></br>
						<br></br>
						<br></br>
						
						<label>Are your current dogs/cats spayed and/or neutered?
						</label>
						<input type="radio" name="current_spayed_neutered" value="Yes"/> 
						<label>Yes</label>
						<input type="radio" name="current_spayed_neutered" value="No"/> 
						<label>No</label>
						<input type="radio" name="current_spayed_neutered" value="N/A"/> 
						<label>N/A</label>
						<br></br>
						<br></br>

						<label >Do you purchase heartworm prevention from a vet? 
						</label>
						<input type="radio" name="heartworm_purchase" value="Yes" /> 
						<label>Yes</label>
						<input type="radio" name="heartworm_purchase" value="No"/> 
						<label>No</label>
						<br></br>
						<br></br>

						<label>If you answered "No" above, 
							please indicate where you purchase prevention from and 
							provide the name brand.</label>
						<input type="textarea" name="no_heartworm"/>
						<br></br>
						<br></br>
						
						<label>Please list any other animals</label>
						<input type="text" name="other_animals" />
						<br></br>
						<br></br>
						
						// Use email format change for vet
						<label> Contact number and name of Vet(s) used. If multiple vets used, please denote which vet
						is associated with with pet</label>
						<input type="textarea" />
						<br></br>
						<br></br>
						
						// do the pet thing again but chnage it to deceased/no longer owned animals 
						<label> List any other previous animals in the past 10 years, 
													Name/Age/Breed/WhatHappened?</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						//same thing here for people in the home
						<label for= "people in home">Name/Age/Relationship (Person 1)</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						<label for= "people in home">Name/Age/Relationship (Person 2)</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						<label for= "people in home">Name/Age/Relationship (Person 3)</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						<label for= "people in home">Name/Age/Relationship (Person 4)</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						<label for= "people in home">List name/age/relationship if there are any other residents in the home.</label>
						<input type="text"/>
						<br></br>
						<br></br>
						
						<label> What type of home do you live in?
							<input type="radio" name= "type_of_house" value="House" required/>
							House
						</label>
						<label>
							<input type="radio" name= "type_of_house" value="Apartment"/>
							Apartment
						</label>
						<label>
							<input type="radio" name= "type_of_house" value="Condo"/>
							Condo
						</label>
						<br></br>
						<br></br>
						
						<label> Do you rent or Own?
							<input type="radio" name="rent_own" value="Rent" required/>
							Rent
						</label> 
						<label>
							<input type="radio" name="rent_own" value="Own"/>
							Own
						</label>
						<br></br>
						<br></br>
						
						<label> If you rent, list landlord's name and phone number<input type="text"/></label>
						<br></br>
						<br></br>
						
						<label>Yes, I give consent to the HSNWLA to contact your veterinarian(s) for records. (Initial Here)
						<input type="text" name="HSNWLA_consent" required /></label>
						<br></br>
						<br></br>
						
						<label> Yes, prior to submitting this application, I have contacted 
						my veterinarian(s) to give them authorization to release records to 
						us (Initial Here)
						<input type="text"/></label>
						<br></br>
						<br></br>
						
						<label> Is someone home during the day?
							<input type="radio" name="person_at_home" value="Yes" required/>
							Yes
						</label>
						<label>
							<input type ="radio" name="person_at_home" value="No"/>
							No 
						</label>
						<br></br>
						<br></br>
						
						<label> If so, who and when?
						<input type="text"/>
						</label>
						<br></br>
						<br></br>
					
						<label> Do you have a fenced yard? <br></br>
							<input type="radio" name="fenced_yard" value="Yes, Chain Link" required/>
							Yes, Chain Link
						</label>
						<br></br>
						<label>
							<input type="radio" name="fenced_yard" value="Yes, Privacy"/>
							Yes, Privacy
						</label>
						<br></br>
						<label>
							<input type="radio" name="fenced_yard" value="Yes, Wireless Electric"/>
							Yes, Wireless Electric
						</label>
						<br></br>
						<label>
							<input type="radio" name="fenced_yard" value="Yes, Other" />
							Yes, Other
						</label>
						<br></br>
						<label>
							<input type="radio" name="fenced_yard" value="Yes, Partial"/>
							Yes, Partial
						</label>
						<br></br>
						<label>
							<input type="radio" name="fenced_yard" value="No"/>
							No
						</label>
						<br></br>
						<br></br>
						
						<label>Where will the dog stay when you are gone for the day?<br></br>
							<input type="radio" name="dog_stay_while_away" value="Inside" reqiured/>
							Inside 
						</label>
						<br></br>
						<label>
							<input type="radio" name="dog_stay_while_away" value="Outside" />
							Outside
						</label>
						<br></br>
						<label>Where will the dogs sleep at night
							<input type="radio" name="where_will_sleep" value=" Inside" required/>
							Inside
						</label>
						<label>
							<input type="radio" name="where_will_sleep" value=" Outside"/>
							Outside
						</label>
						
						<label> What are your thoughts on a dog living outdoors?<input type="textarea" name="thoughts_on_inside_outside"/></label>
						
						
						<label> Where do dogs get heartworm disease from? <input type="textarea"/></label>
						
						<label> Are your dogs current on heartworm prevention?
							<input type="radio"   name="heartworm_prevention" required/>
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
						
						
						<label> Is there anything else you would like us to know about your home? 
						<input type="textarea"/></label>
						
						
						<label> Are you active military?
							<input type="radio" name="active_duty" required/>
							Yes
						</label>
						
						<label>
							<input type="radio" name="active_duty"/>
							No
						</label>
						
						
						<label> If active duty, is there a chance of deployment? 
							<input type="radio"  name="chance_of_deployment"/>
							Yes
						</label>
						
						<label>
							<input type="radio" name="chance_of_deployment"/>
							No
						</label>
						<br></br>
						
						<label>If you are active duty military and deploy, where will your pets go? <input type="text"/> </label>
						<br></br>
						
						<label>
							INCOMPLETE APPLICATIONS WITH MISSING INFORMATION WILL 
							BE RETURNED AS INCOMPLETE AND NOT PROCESSED. PLEASE MAKE SURE 
							ALL FIELDS ARE ACCURATELY FILLED OUT AND YOUR VETERINARIAN HAS 
							BEEN CONTACTED TO RELEASE RECORDS. 
							<input type= "checkbox" required/>
							Yes I, agree 
						</label>
						<br></br>
						
						<label>
						I UNDERSTAND AND AM PREPARED FOR THE FINANCIAL RESPONSIBILITY
						OF OWNING A PET INCLUDING, BUT NOT LIMITED TO ANNUAL VACCINATIONS, AND 
						HEARTWORM PREVENTION. 
						<input type= "checkbox" required/>
						Yes I, agree 
						</label>
						<br></br>
						
						<label>PLEASE NOTE THAT IF YOU SHOULD BECOME UNABLE TO CARE FOR THIS ANIMAL,
						IT MUST BE RETURNED TO THE HSNWLA. IT CAN NOT BE GIVEN AWAY TO ANYONE. 
						(Initial Below) 
						<input type= "checkbox" required/>
						Yes I, agree 
						</label>
						<br></br>
						
						<label> Signature</label>
						<input type="text" required/>
						<br></br>

						<button type="submit">Submit</button>
					
	</form>
    )
}