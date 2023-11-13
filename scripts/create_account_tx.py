import random
import csv
from faker import Faker
from typing import List, AnyStr

fake = Faker()

def main():
    file_headers = ['Id', 'Date', 'Transaction']
    faked_data = create_fake_data(get_quantity())

    with open('account_details.csv', 'w') as opened_file:
        writer = csv.writer(opened_file)
        writer.writerow(file_headers)
        writer.writerows(faked_data)

def get_quantity():
    import os

    env_quantity = os.environ.get('QUANTITY_OF_DATA_TO_GENERATE')
    default_quantity = 4

    try:
        return int(env_quantity) if env_quantity is not None else default_quantity
    except ValueError:
        return default_quantity


def create_fake_data(quantity: int) -> List:
    data = []
    for index in range(quantity):
        temp = []
        temp.append(index)
        temp.append(genearte_random_date_with_format())
        temp.append(generate_random_tx_quantity())
        data.append(temp)
    return data

def genearte_random_date_with_format() -> AnyStr:
    date = fake.date_between(start_date='-1y', end_date='now')
    month_as_str = str(date.month)
    return f'{month_as_str}/{date.day}'

def generate_random_tx_quantity() -> AnyStr:
    return '{:.2f}'.format(round((random.random()*100)*random.choice([-1, 1]), 2))

if __name__ == '__main__':
    main()
