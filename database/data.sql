CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
);
-- one to one (documents)
-- worksheet 1-6
CREATE TABLE risk_criteria (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    field VARCHAR(255) NOT NULL,
    impact_area VARCHAR(255) NOT NULL,
    low TEXT NOT NULL,
    moderate TEXT NOT NULL,
    high TEXT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
-- one to one (documents)
-- worksheet 7
CREATE TABLE impact_priorities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT,
    reputation_confidence INT NOT NULL,
    financial INT NOT NULL,
    productivity INT NOT NULL,
    safety_health INT NOT NULL,
    fines_legal_penalties INT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
-- many to _ (documents)
-- worksheet 8
CREATE TABLE asset_profiles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    critical_asset VARCHAR(255) NOT NULL,
    rationale_for_selection TEXT NOT NULL,
    description TEXT NOT NULL,
    owners TEXT NOT NULL,
    confidentiality TEXT NOT NULL,
    integrity TEXT NOT NULL,
    availability TEXT NOT NULL,
    most_important_security_requirement VARCHAR(255) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
-- many to one (critical_information_asset_profile)
-- worksheet 9
CREATE TABLE asset_containers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT NOT NULL,
    project_id INT NOT NULL,
    owners TEXT NOT NULL,
    technical TEXT,
    physical TEXT,
    people TEXT,
    FOREIGN KEY (asset_id) REFERENCES critical_information_asset_profile(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);
-- many to (documents)
-- worksheet 10
CREATE TABLE asset_risks(
    id INT AUTO_INCREMENT PRIMARY KEY,
    container_id INT NOT NULL,
    project_id INT NOT NULL,
    area_of_concern TEXT NOT NULL,
    actor TEXT NOT NULL,
    means TEXT NOT NULL,
    motive TEXT NOT NULL,
    outcome TEXT NOT NULL,
    security_requirements TEXT NOT NULL,
    probability ENUM('high', 'medium', 'low') NOT NULL,
    consequences TEXT NOT NULL,
    severity ENUM('high', 'medium', 'low') NOT NULL,
    FOREIGN KEY (container_id) REFERENCES critical_asset_containers(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);