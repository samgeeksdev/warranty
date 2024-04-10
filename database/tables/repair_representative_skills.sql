CREATE TABLE repair_representative_skills (
                                              repair_representative_id BIGINT NOT NULL REFERENCES repair_representatives(id),
                                              skill_id BIGINT NOT NULL REFERENCES skills(id),
                                              PRIMARY KEY (repair_representative_id, skill_id)
);
