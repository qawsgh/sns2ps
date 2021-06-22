#!/usr/bin/env python3 

import json
import jinja2
import random
import sys


CATEGORIES = {
           "-": "NoCategory",
           "S": "Senior",
           "SS": "SuperSenior",
           "J": "Junior",
           "SJ": "SuperJunior",
           "L": "Lady"
           }

SG_DIVISIONS = {
                "sg1": "Open",
                "sg2": "Modified",
                "sg3": "Standard",
                "sg4": "StandardManual"
}

AI_DIVISIONS = {
                "ai1": "Open",
                "ai2": "Standard",
                "ai3": "Production",
                "ai3a": "ProductionOptions",
                "ai8": "Classic"
                }

MR_DIVISIONS = {
                "mr1": "MiniRifleOpen",
                "mr2": "MiniRifleStandard",
                "mrc": "Custom"
                }

REGIONS = ["BEL", "DEU", "FRA", "GBR"]

SQUADS = ["27462", "27463", "27464", "27465", "27466", "27467", "27468", "27470", "28842",
          "32408", "29748"]


def generate_ai_competitors_json():
    number = 1
    competitors = []

    templateLoader = jinja2.FileSystemLoader(searchpath="./")
    templateEnv = jinja2.Environment(loader=templateLoader)
    TEMPLATE_FILE = "competitor_json.j2"
    template = templateEnv.get_template(TEMPLATE_FILE)

    for div in AI_DIVISIONS.keys():
        for cat in CATEGORIES.keys():
            sex = 'm'
            if cat == 'L':
                sex = 'f'
            template_vars = {
                "division": AI_DIVISIONS[div],
                "sex": sex,
                "ai_division": div,
                "category": CATEGORIES[cat],
                "category_code": cat,
                "region": random.choice(REGIONS),
                "squad": random.choice(SQUADS),
                "event": "7486",
                "number": number
            }
            number += 1
            competitor = json.loads(template.render(template_vars))
            competitors.append(competitor)

    with open('../jsoncontent/ai_competitors.json', 'w') as f:
        json.dump(competitors, f, indent=2)
    return competitors


def generate_mr_competitors_json():
    number = 1
    competitors = []

    template_loader = jinja2.FileSystemLoader(searchpath="./")
    template_env = jinja2.Environment(loader=template_loader)
    TEMPLATE_FILE = "competitor_json.j2"
    template = template_env.get_template(TEMPLATE_FILE)

    for div in MR_DIVISIONS.keys():
        for cat in CATEGORIES.keys():
            sex = 'm'
            if cat == 'L':
                sex = 'f'
            template_vars = {
                "division": MR_DIVISIONS[div],
                "sex": sex,
                "mr_division": div,
                "category": CATEGORIES[cat],
                "category_code": cat,
                "region": random.choice(REGIONS),
                "squad": random.choice(SQUADS),
                "event": "7486",
                "number": number
            }
            number += 1
            competitor = json.loads(template.render(template_vars))
            competitors.append(competitor)

    with open('../jsoncontent/mr_competitors.json', 'w') as f:
        json.dump(competitors, f, indent=2)
    return competitors


def generate_sg_competitors_json():
    number = 1
    competitors = []

    template_loader = jinja2.FileSystemLoader(searchpath="./")
    template_env = jinja2.Environment(loader=template_loader)
    TEMPLATE_FILE = "competitor_json.j2"
    template = template_env.get_template(TEMPLATE_FILE)

    for div in SG_DIVISIONS.keys():
        for cat in CATEGORIES.keys():
            sex = 'm'
            if cat == 'L':
                sex = 'f'
            template_vars = {
                "division": SG_DIVISIONS[div],
                "sex": sex,
                "sg_division": div,
                "category": CATEGORIES[cat],
                "category_code": cat,
                "region": random.choice(REGIONS),
                "squad": random.choice(SQUADS),
                "event": "7486",
                "number": number
            }
            number += 1
            competitor = json.loads(template.render(template_vars))
            competitors.append(competitor)

    with open('../jsoncontent/sg_competitors.json', 'w') as f:
        json.dump(competitors, f, indent=2)
    return competitors

if __name__ == '__main__':
    generate_ai_competitors_json()
    generate_mr_competitors_json()
    generate_sg_competitors_json()
    

