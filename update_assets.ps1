param(
    [Parameter()]
    [string]$v
)

# Invoke-WebRequest -Uri "https://raw.githubusercontent.com/grindinggear/skilltree-export/$v.0/data.json" -OutFile ./data/SkillTree.json.gz

Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/raw/AlternatePassiveAdditions.json.gz" -OutFile ./data/alternate_passive_additions.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/raw/AlternatePassiveSkills.json.gz" -OutFile ./data/alternate_passive_skills.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/raw/AlternateTreeVersions.json.gz" -OutFile ./data/alternate_tree_versions.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/raw/PassiveSkills.json.gz" -OutFile ./data/passive_skills.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/raw/Stats.json.gz" -OutFile ./data/stats.json.gz

Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/stat_translations/tw/stat_descriptions.json.gz" -OutFile ./data/stat_descriptions.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/stat_translations/tw/passive_skill_stat_descriptions.json.gz" -OutFile ./data/passive_skill_stat_descriptions.json.gz
Invoke-WebRequest -Uri "https://go-pob-data.pages.dev/data/$v/stat_translations/tw/passive_skill_aura_stat_descriptions.json.gz" -OutFile ./data/passive_skill_aura_stat_descriptions.json.gz

go generate -tags tools -x ./...