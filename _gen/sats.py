import itertools
import random
from itertools import product
import faker

orb = ["LEO", "MEO"]
region = ["EU", "AM", "AS", "AU"]
district = ["D%02d"%(i) for i in range(1,31)]
const = ["SATCOM", "STARLINK", "VIASAT", "IRIDIUM"]

map_region = {
    "EU": "Europole",
    "AM": "Amerique",
    "AS": "Asie",
    "AU": "Australie",
}

geo = product(["GEO"], region, district)
autre = product(orb, const)

for g in geo:
    id = "-".join(g)
    keywords = ",".join(['"%s"'%(i) for i in g])
    title = "%s %s"%(map_region[g[1]], g[2])
    region = map_region[g[1]]
    district = "District %s"%(g[2])
    val = "azimut:%f\naltitude:%f\n%s %s"%(random.random()*360, random.random()*90, region, district)
    print('{"%s", []string{%s}, 3, "", "%s", `%s`},'%(id, keywords, title, val))
    
f = faker.Faker()

for c in autre:
    id = "-".join(c)
    keywords = ",".join(['"%s"'%(i) for i in c])
    title = "Constellation %s orbite %s"%(c[1], c[0])
    val = f.uuid4()
    print('{"%s", []string{%s}, 3, "", "%s", "%s"},'%(id, keywords, title, val))
