import names, random, time

# scenario1 scenario
class scenario1:
    global scenario_one_story
    # gym scenaio
    def gym_scenario1():
        # ----------------------------------------------------------------------------#
        # friend system
        friendOne = names.get_first_name()
        friendTwo = names.get_first_name()
        # protag name
        friendZero = names.get_first_name()

    # --------------------------------------------------------------------------#
    # gym chooser
    gym_location = {
        "lol1": "gym near university campus",
        "lol2": "gym near work",
        "lol3": "gym near home",
        "lol4": "home",
    }
    # weekdays
    weekday = {
        "Monday": "Leg day",
        "Tuesday": "Rest day",
        "Wednesday": "Chest day",
        "Thursday": "Rest day",
        "Friday": "Back day",
        "Saturday": "Rest day",
        "Sunday": "Arm days",
    }
    # days and gym co-ordination
    if weekday == "Tuesday" or "Thursday" or "Saturday":
        gym_location = gym_location["lol4"]  # rest days
    else:
        weekday == "Wednesday" or "Friday" or "Sunday" or "Monday"
        gym_location = (
            gym_location["lol1"] or gym_location["lol2"] or gym_location["lol3"]
        )
    # ----------------------------------------------------------------#
    # outfits
    gym_shirt = {
        "shirt1": "red sweatshirt",
        "shirt2": "black sweatshirt",
    }
    gym_pants = {
        "pants1": "black seatpants",
        "pants2": "grey seatpants",
    }
    gym_shoes = {"shoe1": "black tracker shoes", "shoe2": "grey nike shoes"}
    gym_accessories = {"acc1": "black wristwatch", "acc2" : "black smartwatch"}
    # gym outfit
    gym_outfit = (
        gym_shirt["shirt1"]
        or gym_shirt["shirt2"] + gym_pants["pants1"]
        or gym_pants["pants2"] + gym_shoes["shoes1"]
        or gym_shoes["shoes2"] + gym_accessories["accessories1"]
        or gym_accessories["accessories2"]
    )
    print = gym_outfit
